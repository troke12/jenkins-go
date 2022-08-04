pipeline {
    agent { docker { image 'golang:latest' } }
    stages {
        stage('access first') {
            steps {
                sh 'su - -'
            }
        }
        stage('test version') {
            steps {
                sh 'go version'
            }
        }
        stage('go test') {
            steps {
                sh 'go test'
            }
        }
        stage('deploy') {
            steps {
                sh 'docker rmi -f docker-go'
                sh 'docker container prune -f'
                sh 'docker build . -t docker-go'
            }
        }
        stage('start') {
            steps {
                sh 'docker run -d -p 4545:8080 --name docker-go docker-go'
            }
        }
        stage('stop') {
            steps {
                sh 'sleep 120'
                sh 'docker stop docker-go'
            }
        }
    }
}