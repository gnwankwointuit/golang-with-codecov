// pipeline {
//     // Run on an agent where we want to use Go
//     agent any
//
//     // Ensure the desired Go version is installed for all stages,
//     // using the name defined in the Global Tool Configuration
//     tools { go 'go1.20.5' }
//
//     environment {
//         GO111MODULE = 'on'
//         CGO_ENABLED = 1
//         GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
//         GOOS= 'linux'
//         GOARCH='amd64'
//     }
//     stages {
//         stage('Checkout SCM') {
//             steps {
//                 // checkout code from repo
//                 checkout scm
//             }
//         }
//         stage('Test') {
//             steps {
//                 sh 'go version'
//                 sh 'go test -race -coverprofile=coverage.txt -covermode=atomic -v -mod=vendor ./...'
//             }
//         }
//     }
//     post {
//         always {
//             emailext body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}",
//                 recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
//                 to: "${params.RECIPIENTS}",
//                 subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME}"
//
//         }
//     }
// }
 pipeline {
        agent any
        tools { go 'go1.20.5' }

        environment {
            GO111MODULE = 'on'
            CGO_ENABLED = 1
            GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
            GOOS= 'linux'
            GOARCH='amd64'
        }
        stages {
            stage('checkout') {
                steps {
                  checkout scm
                }
            }
            stage('Test') {
                steps {
                  sh 'go test -race -coverprofile=coverage.txt -covermode=atomic -v -mod=vendor ./...'
                }
            }
        }
}