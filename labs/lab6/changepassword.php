<?php
	$username= $_REQUEST["username"];
	$newpassword = $_REQUEST["newpassword"];
	if (isset($username) AND isset($newpassword)) {
		echo "DEBUG:changepassword.php>>Got: username=$username;newpassword=$newpassword\n<br>";
	} else {
		echo "No provided username/password to change.";
		exit();
	}
?>
<a href="index.php">Home</a> | <a href="logout.php">Logout</a>