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
        // Using SSH pipeline steps
        // https://plugins.jenkins.io/ssh-steps/
        stage('Pull to the server and start') {
            when { branch 'master' }
            steps {
                script {
                     withCredentials([string(credentialsId: 'dockerPass', variable: 'dockerpass'), string(credentialsId: 'dockerId', variable: 'dockeruser'), string(credentialsId: 'servDev', variable: 'servDev'), string(credentialsId: 'sshId', variable: 'sshUser'), string(credentialsId: 'sshPw', variable: 'sshpass')]) {
                        //Dev server
                        def dev = [:]
                        dev.name = 'Development Server'
                        dev.host = servDev
                        dev.user = sshUser
                        dev.password = sshpass
                        dev.allowAnyHosts = true
                        //sshCommand remote: dev, command: 'ls -l'
                        sshCommand remote: dev, command: "docker login -u ${dockeruser} -p ${dockerpass} registry.indoteam.id"
                        sshCommand remote: dev, command: "docker pull registry.indoteam.id/indoteam/jenkins-go-master:${BUILD_NUMBER}"
                        sshCommand remote: dev, command: "docker run -d -p 6060:8080 --name jenkins-go-master registry.indoteam.id/indoteam/jenkins-go-master:${BUILD_NUMBER}"
                    }
                }
            }
        }
    }
}