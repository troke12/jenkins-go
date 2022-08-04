pipeline {
    agent any

    stages {
        stage('Build & Test') {   
            // Use golang.
            agent { docker { image 'golang' } }

            steps {                                           
                // Create our project directory.
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/jenkins-go'

                // Copy all files in our Jenkins workspace to our project directory.                
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/jenkins-go'

                // Build the app.
                sh 'go build'

                // Clean cache
                sh 'go clean --cache'               
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