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
                withCredentials([string(credentialsId: 'dockerPass', variable: 'dockerpass'), string(credentialsId: 'dockerId', variable: 'dockeruser')]) {
                    sh 'docker login -u $dockeruser -p $dockerpass registry.indoteam.id'
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