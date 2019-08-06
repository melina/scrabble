pipeline {
    agent any
    
    tools {
        go 'go-1.11'
    }
    environment {
        GO111MODULE = 'on'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Test') {
            steps{
                sh """
                go test -v
                """
            }
            
        }

        stage('Push image') {
            agent { label 'agent' }
            steps {
                echo 'Starting to build docker image'
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
                        script {
                            def app = docker.build("melibou/scrabble:${env.BUILD_ID}")
                            app.push("${env.BUILD_NUMBER}")
                            app.push("latest")
                        }
                    }
                } 
            }
        }
    }
}
