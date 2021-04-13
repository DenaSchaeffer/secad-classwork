<?php
	$lifetime = 15 *60; //15 minutes
	$path = "/lab6";
	$domain ="192.168.210.3"; //own ip address
	$secure = TRUE;
	$httponly = TRUE;
	session_set_cookie_params($lifetime, $path, $domain, $secure, $httponly);
	session_start();  
	$mysqli = new mysqli('localhost',
							'backd1', /*Database username*/
							's3cur3p4ss', /*Database password*/
							'secad'); /*Database name*/
		if($mysqli->connect_errno) {
			printf("Database connection failed: %s\n", $mysqli->connect_error);
			exit();
		}
	if(isset($_POST["username"]) and isset($_POST["password"])) {

		if (securechecklogin($_POST["username"],$_POST["password"])) {
			$_SESSION["logged"] = TRUE;
			$_SESSION["username"] = $_POST["username"];
			$_SESSION["browser"] = $_SERVER["HTTP_USER_AGENT"];
		}else{
			echo "<script>alert('Invalid username/password');</script>";
			session.destroy();
			header("Refresh:0; url=form.php");
			die();
		}
	}  
	if ($_SESSION["browser"] != $_SERVER["HTTP_USER_AGENT"]) {
		echo "<script>alert('Session hijacking is detected!');</script>";
		header("Refresh:0; url=form.php");
		die();
	}
	if (!isset($_SESSION["logged"]) or $_SESSION["logged"] != TRUE) {
		echo "<script>alert('You have not logged in. Please login first');</script>";
		header("Refresh:0; url=form.php");
		die();
	}
		
?>
	<h2> Welcome <?php echo htmlentities($_SESSION['username']); ?> !</h2>
	<a href="logout.php">Logout</a>
	<a href="changepasswordform.php">Change password</a>
<?php	
  	function securechecklogin($username, $password){
  		global $mysqli;
  		$prepared_sql = "SELECT * FROM users WHERE username= ? " .
  						" AND password = password(?);";
		//echo "DEBUG>sql= $sql"; //return TRUE

  		if(!$stmt = $mysqli->prepare($prepared_sql))
  			echo "Prepared Statement Error";
  		$stmt->bind_param('ss',$username,$password);
  		if (!$stmt->execute())
  			echo "Execute Error";
  		if (!$stmt->store_result())
  			echo "Store Result Error";
  		$result = $stmt; //$mysqli->query($sql); //send the SQL query to the database
  		if($result->num_rows ==1)
  			return TRUE;
  		return FALSE;
  	}
?>
