# /bin/sh

mvn -f pom2.xml clean install -Dmaven.test.skip=true -Ptest -e

cd $1
mvn clean package -Dmaven.test.skip=true -Ptest -e