pipeline {
    agent {
        dockerfile true
    }
    environment {
            GOROOT = "${env.WORKSPACE}/go"
            GOPATH = "${env.WORKSPACE}/gopath"
            PATH = "${env.PATH}:${GOROOT}/bin:${GOPATH}/bin"
        }
    stages {
        stage('checkout') {
            steps {
                sh 'git checkout main'
            }
        }
        stage('test') {
            steps {
                sh 'pwd'
                sh ''
                sh '/bin/bash -c "go build ./..."'
                sh '/bin/bash -c "go test ./..."'
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
                sh 'ls -al ${env.WORKSPACE}'
                deleteDir()
                sh 'ls -al ${env.WORKSPACE}'
            }
        }
    }
}
