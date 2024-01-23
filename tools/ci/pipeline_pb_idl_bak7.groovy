def CHANGE_LIST_ID = ''
def BRANCH_NAME = 'pb_idl_auto_generate'
def PR_TITLE = 'pb_idl Files (Commit %s)'

def PR_BODY = """
This pull request updates the auto-generated files for the pb_idl submodule based on the latest changes.

### Commit Details
- **Commit ID**: %s
- **Summary**: Auto-generated updates for pb_idl files.

Please review the changes and merge them into the main branch if everything looks good.
"""


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
                    /// Checkout the openiiot repository
                    dir('pb_idl') {
                        checkout scmGit(branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/hongzhi1234/pb_idl.git']])
                        //checkout([$class: 'GitSCM', branches: [[name: '*/main']], userRemoteConfigs: [[url: 'https://github.com/hongzhi1234/openiiot.git']]])
                    }

                    sh 'git clone https://github.com/freezonex/openiiot.git'
                    sh 'cd openiiot && git submodule update --init --recursive pb_idl'
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
                    def currentDir = sh(script: 'pwd', returnStdout: true).trim()
                    echo "Current directory is: ${currentDir}"

                    if (CHANGE_LIST_ID != '') {
                        //BRANCH_NAME = "pb_idl_${CHANGE_LIST_ID}"
                        sh "cd openiiot && git checkout -b ${BRANCH_NAME}"
                        sh 'cd openiiot && hz model --idl pb_idl/freezonex/openiiot_api/openiiot_api_service.proto --unset_omitempty -I pb_idl'
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
                        PR_TITLE = String.format(PR_TITLE, CHANGE_LIST_ID)
                        PR_BODY = String.format(PR_BODY, CHANGE_LIST_ID)
                        withCredentials([string(credentialsId: 'hongzhi_token', variable: 'GITHUB_TOKEN')]) {
                            dir('openiiot') {
                                sh """
                                    git config user.email 'zhanghongzhi@freezonex.io'
                                    git config user.name 'hongzhi1234'
                                    git remote set-url origin https://hongzhi1234:${GITHUB_TOKEN}@github.com/hongzhi1234/openiiot.git
                                    git add biz/model pb_idl
                                    git commit -m 'Update files created by hz tool'
                                """
                                
                                sh 'git fetch --prune'  // Fetch the latest branch information
                                //def branchExists = sh(script: "git branch -r | grep -w 'origin/${BRANCH_NAME}'", returnStdout: true).trim()  // Check if the branch exists on the remote
                                def branchExists = sh(script: "git branch -r | grep -w 'origin/${BRANCH_NAME}' || true", returnStdout: true).trim()  // Check if the branch exists on the remote

                                if (branchExists) {
                                    sh "git push origin --delete ${BRANCH_NAME}"     // Delete the branch if it exists
                                }

                                sh """
                                    git push origin ${BRANCH_NAME}
                                    curl -X POST -H 'Authorization: token ${GITHUB_TOKEN}' -d '{\"title\": \"${PR_TITLE}\", \"body\": \"${PR_BODY}\", \"head\": \"${BRANCH_NAME}\", \"base\": \"main\"}' https://api.github.com/repos/hongzhi1234/openiiot/pulls
                                """
                            }
                        }
                    } else {
                        echo 'Skipping "Create Pull Request" stage due to no new changelist.'
                    }
                }
            }
        }
    }
}
