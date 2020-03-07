FROM tutum/lamp:latest

ENV DEBIAN_FRONTEND noninteractive

# Preparation
RUN rm -fr /app/* && \
  apt-get update && apt-get install -yqq wget unzip php5-curl dnsutils && \
  rm -rf /var/lib/apt/lists/*

# Deploy Mutillidae
RUN \
  wget -O /mutillidae.zip https://jaist.dl.sourceforge.net/project/mutillidae/mutillidae-project/LATEST-mutillidae-2.6.62.zip && \
  unzip /mutillidae.zip && \
  rm -rf /app/* && \
  cp -r /mutillidae/* /app  && \
  rm -rf /mutillidae  && \
  sed -i 's/DirectoryIndex index.html.*/DirectoryIndex index.php index.html index.cgi index.pl index.xhtml index.htm/g' /etc/apache2/mods-enabled/dir.conf&& \
  sed -i 's/static public \$mMySQLDatabaseUsername =.*/static public \$mMySQLDatabaseUsername = "admin";/g' /app/classes/MySQLHandler.php && \
  echo "sed -i 's/static public \$mMySQLDatabasePassword =.*/static public \$mMySQLDatabasePassword = \\\"'\$PASS'\\\";/g' /app/classes/MySQLHandler.php" >> /create_mysql_admin_user.sh && \
  echo 'session.save_path = "/tmp"' >> /etc/php5/apache2/php.ini 

EXPOSE 80 3306
CMD ["/run.sh"]
