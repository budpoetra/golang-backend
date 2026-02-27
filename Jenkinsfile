pipeline {
    agent any
    
    triggers {
        pollSCM('* * * * *')
    }

    environment {
        APP_NAME = "golang-backend"
        GO_VERSION = "1.21"
    }

    stages {
        stage('Hello') {
            steps {
                echo("Hello World!")
            }
        }
        
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        // stage('Build & Test') {
        //     steps {
        //         sh 'go mod tidy'
        //         sh 'go test ./...'
        //         sh "go build -o ${APP_NAME}"
        //     }
        // }
    }
}
