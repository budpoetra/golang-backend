pipeline {
    agent any
    
    triggers {
        pollSCM('* * * * *')
    }

    tools {
        go 'Go-1.26.0'
    }

    environment {
        APP_NAME = "golang-backend"
        GO_VERSION = "1.21"
    }

    options {
        buildDiscarder(logRotator(
            numToKeepStr: '10',
            daysToKeepStr: '30',
            artifactNumToKeepStr: '5',
            artifactDaysToKeepStr: '15'
        ))
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
        
        stage('Build & Test') {
            steps {
                sh 'go mod tidy'
                sh 'go test ./...'
                sh "go build -o ${APP_NAME}"
            }
        }
    }
}
