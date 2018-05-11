pipeline {
  agent none
  stages {
    stage('Test') {
      agent {
        docker {
        image 'golang'
        }
      }
      steps {
        sh 'go test'
      }
    }
  stage('Deploy') {
    agent none
    steps {
      script {
        def customImage = docker.build("bdubois/maths:${env.BUILD_NUMBER}")
        docker.withRegistry('https://registry.hub.docker.com', 'Docker') {
            customImage.push("${env.BUILD_NUMBER}")
            customImage.push("latest")
        }
      }
    }
  }
}
}