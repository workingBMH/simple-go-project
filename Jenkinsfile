pipeline {
    agent any
    stages {
        stage('checkout') {
            steps {
                sh 'git checkout main'
            }
        }
        stage('test and build') {
            steps {
                sh 'go test ./...'
                
            }
            when {
                expression {
                    expressionBuild.result == null || currentBuild.result == 'SUCCESS'
                }
            }
            steps{
                sh 'go build ./...'
            }

        }
        stage('push to hub') {
            when {
                expression {
                    expressionBuild.result == null || currentBuild.result == 'SUCCESS'
                }
            }
            steps {
                sh 'docker build -t ne275/simple-go-project-jenkins:jenkinsfile -f Dockerfile .'
                sh 'docker push ne275/simple-go-project-jenkins:jenkinsfile'
            }

        }
        stage('clean local images') {
            steps { 
                sh 'docker rmi ne275/simple-go-project-jenkins:jenkinsfile'
                sh 'docker builder prune'
            }
            steps {
                sh 'ls -al ${env.WORKSPACE}'
                deleteDir()
                sh 'ls -al ${env.WORKSPACE}'
            }

        }
    }
}