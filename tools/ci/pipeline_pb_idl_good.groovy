def CHANGE_LIST_ID = ''
def BRANCH_NAME = ''

pipeline {
    options {
        skipStagesAfterUnstable()
    }
    agent { label 'jenkins-golang-agent' } // Specify the label of your Jenkins agent
    stages {
        stage('Clean Workspace') {
            steps {
                deleteDir() // Deletes the workspace's contents
            }
        }

        stage('Prepare and Check for pb_idl main branch') {
            steps {
                script {
                    // Clone the openiiot repo
                    sh 'git clone https://github.com/hongzhi1234/openiiot.git'

                    // Initialize and update the submodule
                    sh 'cd openiiot && git submodule update --init --recursive pb_idl'

                    // Update pb_idl submodule to the latest commit on its main branch
                    sh 'cd openiiot/pb_idl && git checkout main && git pull'
                    
                    // Check the submodule status
                    def submoduleStatus = sh(script: "cd openiiot && git submodule status pb_idl", returnStdout: true).trim()
                    echo "Submodule Status: ${submoduleStatus}"


                    // Try to extract the commit hash using regex
                    def match = (submoduleStatus =~ /\+(\w+)/)
                    if (match) {
                        def commitHash = match[0][1]
                        echo "Regex Match Result: ${commitHash}"

                        CHANGE_LIST_ID = commitHash
                        echo "pb_idl submodule has new updates on main branch. Changelist ID: ${CHANGE_LIST_ID}"
                    } else {
                        // Handle the case where there is no '+' sign (i.e., no new changelist)
                        echo "No new updates in pb_idl submodule. Stopping the pipeline."
                        //error("No new updates - stopping pipeline")
                        currentBuild.result = 'SUCCESS'
                        return
                    }
                }
            }
        }

        stage('Run hz command') {
            steps {
                script {
                    if (CHANGE_LIST_ID != '') {
                        BRANCH_NAME = "pb_idl_${CHANGE_LIST_ID}"
                        dir('openiiot') {
                            sh "git checkout -b ${BRANCH_NAME}"
                            sh 'hz model --idl pb_idl/freezonex/openiiot_api/openiiot_api_service.proto --unset_omitempty -I pb_idl'
                        }
                    } else {
                        echo 'Skipping "Run hz command" stage due to no new changelist.'
                    }
                }
            }
        }

        stage('Create Pull Request') {
            steps {
                script {
                    if (CHANGE_LIST_ID != '') {
                        withCredentials([string(credentialsId: 'hongzhi_token', variable: 'GITHUB_TOKEN')]) {
                            sh """
                                cd openiiot
                                git config user.email 'zhanghongzhi@freezonex.io'
                                git config user.name 'hongzhi1234'
                                git remote set-url origin https://hongzhi1234:${GITHUB_TOKEN}@github.com/hongzhi1234/openiiot.git
                                git add *
                                git commit -m 'Update files created by hz tool'
                                git push origin ${BRANCH_NAME}
                                curl -X POST -H 'Authorization: token ${GITHUB_TOKEN}' -d '{\"title\": \"pb_idl update\", \"body\": \"Update auto-generated files for changelist: ${CHANGE_LIST_ID}\", \"head\": \"${BRANCH_NAME}\", \"base\": \"main\"}' https://api.github.com/repos/hongzhi1234/openiiot/pulls
                            """
                        }
                    } else {
                        echo 'Skipping "Create Pull Request" stage due to no new changelist.'
                    }
                }
            }
        }
    }
}
