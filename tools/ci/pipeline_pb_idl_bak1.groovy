pipeline {
    agent any
    environment {
        // Define a variable to store the branch check result
        IS_MAIN_BRANCH = 'false'
    }

    stages {
        stage('Check for pb_idl main branch') {
            steps {
                script {
                    // Perform the branch check and update the variable
                    IS_MAIN_BRANCH = sh(script: "cd openiiot && git submodule status pb_idl | grep 'origin/hongzhi1'", returnStdout: true).trim()
                }
            }
        }

        stage('Update openiiot') {
            when {
                // Use the variable in the expression
                expression { IS_MAIN_BRANCH == 'true' }
            }
            steps {
                script {
                    // Clone the openiiot repo
                    sh 'git clone https://github.com/freezonex/openiiot.git'
                    dir('openiiot') {
                        // Update the pb_idl submodule
                        sh 'git submodule update --init pb_idl'
                    }
                }
            }
        }

        stage('Run Interface Update Command') {
            when {
                expression { IS_MAIN_BRANCH == 'true' }
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
