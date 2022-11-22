<?php

$request_headers = apache_request_headers();
echo $request_headers['Host'];

?>
