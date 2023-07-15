pipeline {
	environment {
		TS = sh(script: 'date +%Y-%m-%d-%H-%M', returnStdout: true).trim()
	}

	agent any

	stages {
		stage('checkout') {
			steps {
				sh "echo ${STAGE_NAME}"
			}
			post {
				aborted {
					sh "echo ${STAGE_NAME}: aborted"
				}
				failure {
					sh "echo ${STAGE_NAME}: failure"
				}
				success {
					sh "echo ${STAGE_NAME}: success"
				}
			}
		}

		stage('build') {
			steps {
				sh "echo ${STAGE_NAME}"
			}
			post {
				aborted {
					sh "echo ${STAGE_NAME}: aborted"
				}
				failure {
					sh "echo ${STAGE_NAME}: failure"
				}
				success {
					sh "echo ${STAGE_NAME}: success"
				}
			}
		}

		stage('test') {
			steps {
				sh "echo ${STAGE_NAME}"
			}
			post {
				aborted {
					sh "echo ${STAGE_NAME}: aborted"
				}
				failure {
					sh "echo ${STAGE_NAME}: failure"
				}
				success {
					sh "echo ${STAGE_NAME}: success"
				}
			}
		}

		stage('publish') {
			steps {
				sh "echo ${STAGE_NAME}"
			}
			post {
				aborted {
					sh "echo ${STAGE_NAME}: aborted"
				}
				failure {
					sh "echo ${STAGE_NAME}: failure"
				}
				success {
					sh "echo ${STAGE_NAME}: success"
				}
			}
		}

	}

	post {
		always {
			cleanWs()
		}
	}
}
