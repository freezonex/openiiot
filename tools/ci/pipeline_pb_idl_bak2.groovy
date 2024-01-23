pipeline {
    agent any

    stages {
        stage('Clean Workspace') {
            steps {
                deleteDir() // This deletes the workspace's contents
            }
        }

        stage('Prepare and Check for pb_idl main branch') {
            steps {
                script {
                    // Clone the openiiot repo
                    sh 'git clone https://github.com/hongzhi1234/openiiot.git'
                    // Initialize and update the submodule
                    sh 'cd openiiot && git submodule update --init --recursive'

                    // Get the latest commit hash of the pb_idl main branch
                    def latestCommit = sh(script: "cd openiiot/pb_idl && git fetch && git rev-parse origin/main", returnStdout: true).trim()

                    // Get the current commit hash of the pb_idl submodule
                    def currentCommit = sh(script: "cd openiiot && git submodule status pb_idl | awk '{print \$1}'", returnStdout: true).trim()

                    // Compare the hashes
                    if (currentCommit != latestCommit) {
                        error("pb_idl submodule is not at the latest commit of main branch. Skipping further steps.")
                    }
                }
            }
        }

        stage('Update openiiot') {
            when {
                expression { return isMainBranch }
            }
            steps {
                script {
                    dir('openiiot') {
                        // Update the pb_idl submodule
                        sh 'git submodule update --init pb_idl'
                    }
                }
            }
        }

        stage('Install Dependencies') {
            steps {
                script {
                    // Install hz
                    sh 'go install github.com/cloudwego/hertz/cmd/hz@latest'

                    // Install Protocol Buffers (protobuf)
                    sh 'sudo apt-get install -y protobuf-compiler'
                }
            }
        }

        stage('Run Interface Update Command') {
            when {
                expression { return isMainBranch }
            }
            steps {
                script {
                    dir('openiiot') {
                        // Run your specific command
                        sh 'hz model --idl pb_idl/freezonex/openiiot_api/openiiot_api_service.proto --unset_omitempty -I pb_idl'
                    }
                }
            }
        }

        stage('Create Pull Request') {
            when {
                expression { IS_MAIN_BRANCH == 'true' }
            }
            steps {
                script {
                    // Additional scripting for creating a pull request
                    // This might involve using the GitHub API or a Git tool
                }
            }
        }

    }
}
