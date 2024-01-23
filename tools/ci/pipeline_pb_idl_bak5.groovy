pipeline {
    agent any

    stages {
        stage('Clean Workspace') {
            steps {
                deleteDir() // This deletes the workspace's contents
            }
        }

        stage('Prepare Environment') {
            steps {
                script {
                    // Install Golang if it's not installed
                    sh '''
                    if ! type "go" > /dev/null; then
                      sudo apt-get update
                      sudo apt-get install -y golang-go
                    fi
                    go version
                    '''

                    // Install Protocol Buffers Compiler
                    sh 'sudo apt-get install -y protobuf-compiler'
                }
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
                    
                    // Check if the submodule status starts with '+'
                    if (submoduleStatus.startsWith("+")) {
                        echo "pb_idl submodule has new updates on main branch."
                    } else {
                        error("No new updates in pb_idl submodule.")
                    }
                }
            }
        }

        stage('Run hz command') {
            steps {
                script {
                    // Install hz (assuming Go is already installed and configured)
                    sh 'go install github.com/cloudwego/hertz/cmd/hz@latest'

                    // Run hz command
                    sh 'cd openiiot && hz model --idl pb_idl/freezonex/openiiot_api/openiiot_api_service.proto --unset_omitempty -I pb_idl'
                }
            }
        }

        stage('Create Pull Request') {
            steps {
                script {
                    withCredentials([string(credentialsId: 'hongzhi_token', variable: 'GITHUB_TOKEN')]) {
                        // Configure Git to use the PAT for HTTPS operations
                        sh 'git config --global user.email "zhanghongzhi@freezonex.io"'
                        sh 'git config --global user.name "hongzhi1234"'
                        sh 'git config --global http.extraHeader "Authorization: token ${GITHUB_TOKEN}"'

                        // Change directory, add, commit, and push changes
                        sh 'cd openiiot && git add * && git commit -m "Update files created by hz tool" && git push'

                        // Optional: Create a pull request via GitHub API
                        sh '''
                            curl -X POST -H "Authorization: token ${GITHUB_TOKEN}" \
                                -d '{"title":"New PR","head":"your-branch","base":"main"}' \
                                https://api.github.com/repos/hongzhi1234/openiiot/pulls
                        '''
                    }
                }
            }
        }
    }
}

