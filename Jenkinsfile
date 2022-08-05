//Prefixes
def prefix_master = "master"
def prefix_dev = "dev"
def prefix_stag = "stag"
//Start
pipeline {
    agent any
    stages {
        stage('build') {
            steps {
                script {
                    if (env.BRANCH_NAME == "master") {
                        echo 'Starting building for master'
                        sh 'docker build -t registry.indoteam.id/indoteam/jenkins-go-${prefix_master}:${BUILD_NUMBER} .'
                    } if else (env.BRANCH_NAME == "development") {
                        echo 'Starting building for development'
                        sh 'docker build -t registry.indoteam.id/indoteam/jenkins-go-${prefix_dev}:${BUILD_NUMBER} .'
                    } if else (env.BRANCH_NAME == "staging") {
                        echo 'Starting building for staging'
                        sh 'docker build -t registry.indoteam.id/indoteam/jenkins-go-${prefix_stag}:${BUILD_NUMBER} .'
                    }
                }
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
                script {
                    if (env.BRANCH_NAME == "master") {
                        echo 'Starting pushing for master'
                        sh 'docker push registry.indoteam.id/indoteam/jenkins-go-${prefix_master}:${BUILD_NUMBER}'
                    } if else (env.BRANCH_NAME == "development") {
                        echo 'Starting pushing for development'
                        sh 'docker push registry.indoteam.id/indoteam/jenkins-go-${prefix_dev}:${BUILD_NUMBER}'
                    } if else (env.BRANCH_NAME == "staging") {
                        echo 'Starting pushing for staging'
                        sh 'docker push registry.indoteam.id/indoteam/jenkins-go-${prefix_stag}:${BUILD_NUMBER}'
                    }
                }
            }
        }
    }
}