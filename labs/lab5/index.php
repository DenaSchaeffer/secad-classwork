<?php
	session_start();  
	$mysqli = new mysqli('localhost',
							'backd1', /*Database username*/
							's3cur3p4ss', /*Database password*/
							'secad'); /*Database name*/
		if($mysqli->connect_errno) {
			printf("Database connection failed: %s\n", $mysqli->connect_error);
			exit();
		}  
	if (securechecklogin($_POST["username"],$_POST["password"])) {
?>
	<h2> Welcome <?php echo $_POST['username']; ?> !</h2>
<?php		
	}else{
		echo "<script>alert('Invalid username/password');</script>";
		die();
	}
	function checklogin($username, $password) {
		$account = array("admin","1234");
		if (($username== $account[0]) and ($password == $account[1])) 
		  return TRUE;
		else return FALSE;
  	}
  	function checklogin_mysql($username, $password) {
		$mysqli = new mysqli('localhost',
							'backd1', /*Database username*/
							's3cur3p4ss', /*Database password*/
							'secad'); /*Database name*/
		if($mysqli->connect_errno) {
			printf("Database connection failed: %s\n", $mysqli->connect_error);
			exit();
		}
		$sql = "SELECT * FROM users WHERE username='" . $username . "' ";
		$sql = $sql . " AND password = password('" . $password . "')";
		//echo "DEBUG>sql= $sql"; //return TRUE
		$result = $mysqli->query($sql); //send the SQL query to the database
		if($result->num_rows ==1) // if there is record matched
			return TRUE;
		return FALSE;
  	}
  	function securechecklogin($username, $password){
  		global $mysqli;
  		$prepared_sql = "SELECT * FROM users WHERE username= ? AND password = password(?)";
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
