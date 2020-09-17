#!/bin/bash

sed -i "s/p@ssw0rd//g" /app/config/config.inc.php && \
sed -i "s/$_DVWA[ 'recaptcha_private_key' ] = ''/$_DVWA[ 'recaptcha_private_key' ] = '6LdK7xITAAzzAAJQTfL7fu6I-0aPl8KHHieAT_yJg'/g" /app/config/config.inc.php && \
sed -i "s/$_DVWA[ 'recaptcha_public_key' ] = ''/$_DVWA[ 'recaptcha_public_key' ] = '6LdK7xITAzzAAL_uw9YXVUOPoIHPZLfw2K1n5NVQ'/g" /app/config/config.inc.php
