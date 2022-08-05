pipeline {
    agent any
    stages {
        stage('build') {
            steps {
                sh 'docker build -t registry.indoteam.id/indoteam/jenkins-go:${BUILD_NUMBER} .'
            }
        }
        stage('login docker') {
            steps {
                withCredentials([string(credentialsId: 'c979ffa3-2123-4e81-a6a3-db67052f5779', variable: 'docker-password'), string(credentialsId: 'c979ffa3-2123-4e81-a6a3-db67052f5779', variable: 'docker-username')]) {
                    sh 'docker login -u $docker-username -p $docker-password registry.indoteam.id'
                }
            }
        }
        stage('push') {
            steps {
                sh 'docker push registry.indoteam.id/indoteam/jenkins-go:${BUILD_NUMBER}'
            }
        }
    }
}