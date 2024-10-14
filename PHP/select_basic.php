<!--select_basic.php-->
<html>
  <head>
    <title>Operazioni</title>
  </head>
<body>
<?php
print ("<form action=\"opt_basic.php\" method=\"POST\">");
//definisco un selettore statico per le tabelle del db
print ("<select name=\"tabella\">");
//inserisco le tabelle che voglio gestire
print ("<option value=\"cliente\">cliente</option>");
print ("</select>");
//definisco un selettore statico per le operazioni sulle tabelle
print ("<select name=\"operazione\">");
print ("<option value=\"select\">select</option>");
print ("<option value=\"insert\">insert</option>");
print ("<option value=\"update\">update</option>");
print ("<option value=\"delete\">delete</option>");
print ("</select>");
print ("<input  type=\"submit\" >");
print ("</form>");
?>
</body>
</html>