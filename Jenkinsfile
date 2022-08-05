//Start
pipeline {
    agent any
    stages {
        stage('Build For Dev') {
            when { branch 'development' }
            steps {
                echo 'Starting building for Dev'
                sh 'docker build -t registry.indoteam.id/indoteam/jenkins-go-dev:${BUILD_NUMBER} .'
            }
        }
        stage('Build For Staging') {
            when { branch 'staging' }
            steps {
                echo 'Starting building for staging'
                sh 'docker build -t registry.indoteam.id/indoteam/jenkins-go-staging:${BUILD_NUMBER} .'
            }
        }
        stage('Build For Master') {
            when { branch 'master' }
            steps {
                echo 'Starting building for master'
                sh 'docker build -t registry.indoteam.id/indoteam/jenkins-go-master:${BUILD_NUMBER} .'
            }
        }
        stage('login docker') {
            steps {
                withCredentials([string(credentialsId: 'dockerPass', variable: 'dockerpass'), string(credentialsId: 'dockerId', variable: 'dockeruser')]) {
                    sh 'docker login -u $dockeruser -p $dockerpass registry.indoteam.id'
                }
            }
        }
        stage('push for dev') {
            when { branch 'development' }
            steps {
                echo 'Starting pushing for dev'
                sh 'docker push registry.indoteam.id/indoteam/jenkins-go-dev:${BUILD_NUMBER}'
            }
        }
        stage('push for staging') {
            when { branch 'staging' }
            steps {
                echo 'Starting pushing for staging'
                sh 'docker push registry.indoteam.id/indoteam/jenkins-go-staging:${BUILD_NUMBER}'
            }
        }
        stage('push for master') {
            when { branch 'master' }
            steps {
                echo 'Starting pushing for master'
                sh 'docker push registry.indoteam.id/indoteam/jenkins-go-master:${BUILD_NUMBER}'
            }
        }
        stage('clear cache') {
            steps {
                sh 'docker system prune -af'
                echo 'Cached cleared....'
            }
        }
    }
}