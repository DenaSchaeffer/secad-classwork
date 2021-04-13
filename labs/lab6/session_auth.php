
<?php
    $lifetime = 15 * 60;
    $path = "/lab6";
    $domain = "192.168.210.3"; //my ip
    $secure = TRUE;
    $httponly = TRUE;
    session_set_cookie_params($lifetime,$path,$domain,$secure,$httponly);
    session_start();

     //check the session
     if( !isset($_SESSION["logged"]) or $_SESSION["logged"] != TRUE){
     //the session is not authenticated
      echo "<script>alert('You have to login first!');</script>";
      session_destroy();
      header("Refresh:0; url=form.php");
      die();
    }
      if( $_SESSION["browser"] != $_SERVER["HTTP_USER_AGENT"]){
      //it is a session hijacking attack since it comes from a different browser
      echo "<script>alert('Session hijacking attack is detected!');</script>";
      session_destroy();
      header("Refresh:0; url=form.php");
      die();
    }
?>