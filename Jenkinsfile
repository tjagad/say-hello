pipeline {
	environment {
		TS = sh(script: 'date +%Y-%m-%d-%H-%M', returnStdout: true).trim()
	}

	agent any

	stages {
		stage('checkout') {
			steps {
				script {
					checkout changelog: true, poll: true, scm: [
						$class: 'GitSCM',
						branches: [[name: "origin/${env.BRANCH_NAME}"]],
						doGenerateSubmoduleConfigurations: false,
						extensions: [
							[$class: 'GitLFSPull'],
							[$class: 'CheckoutOption', timeout: 60],
							[$class: 'CloneOption', depth: 0, timeout: 60],
						],
						submoduleCfg: [],
						userRemoteConfigs: scm.userRemoteConfigs
					]
				}
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
				sh "echo ${BRANCH_NAME}"
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
				sh "echo ${BRANCH_NAME}"
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
				sh "echo ${BRANCH_NAME}"
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
