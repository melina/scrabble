node {
    def root = tool name: 'Go 1.12.6', type: 'go'
    
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        sh 'go version'
    }

    stage('Clone repository') {
        /* Let's make sure we have the repository cloned to our workspace */

        checkout scm
    }


    stage('Test scrabble') {
        sh 'go test -v'
    }

    stage('Build image') {
        /* This builds the actual image; synonymous to
         * docker build on the command line */

        app = docker.build("melibou/scrabble")
    }


    stage('Push image') {
        /* Finally, we'll push the image with two tags:
         * First, the incremental build number from Jenkins
         * Second, the 'latest' tag.
         * Pushing multiple tags is cheap, as all the layers are reused. */
        docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
            app.push("${env.BUILD_NUMBER}")
            app.push("latest")
        }
    }
}
