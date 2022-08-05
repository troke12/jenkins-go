//Prefixes
def prefix_master = "master"

//Start
pipeline {
    agent any
    stages {
        stage('Build For Production') {
            when { branch 'development' }
            steps {
                echo 'Branch dev so skip'
            }
        }
        stage('Build For Production') {
            when { branch 'master' }
            steps {
                echo 'brahuuuuuuuu master'
            }
        }
        stage('build') {
            steps {
                echo 'Starting building for master'
                sh 'docker build -t registry.indoteam.id/indoteam/jenkins-go-${prefix_master}:${BUILD_NUMBER} .'
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
                echo 'Starting pushing for master'
                sh 'docker push registry.indoteam.id/indoteam/jenkins-go-${prefix_master}:${BUILD_NUMBER}'
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