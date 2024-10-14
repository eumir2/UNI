<?php
    //session_start();
    $htmlint = <<<NOW
    <HTML>
    <HEAD>
        <style>
        table, th, td {
            text-align: left;
            border: 1px solid;
            padding: 5px;
        }
        </style>
    </HEAD>
    <BODY> 
    NOW;
    print ($htmlint);
    $conn = pg_connect("host=localhost port=5432 dbname=Piscine user=postgres password=unimi");
    if (!$conn) {//caso connessione fallita
        echo 'Connessione al database fallita.';
        exit();
    }

    $query = "SELECT * FROM medico";
    $medici = pg_query($conn, $query);
    if (!$medici) {
        echo "Si è verificato un errore.<br/>";
        echo pg_last_error($conn);
        echo ("<br><br>Se vuoi puoi <a href='inserimento_e_modifica.php'>riprovare</a>");
        exit();
    }

    echo '<h1>Inserimento certificato medico</h1>';
    print ("<form action=\"ins.php\" method=\"POST\">");
    echo '<table>';
    //passo le informazioni oggetto dell'inserimento
    print ("<tr><th>Medico</th><td><select name=\"medico\" required>");
    while ($row = pg_fetch_array($medici)){
        echo("<option>".$row['nome']." ".$row['cognome']."</option>");
    }
    echo "</td</tr>";
    print ("<tr><th>Data rilascio</th><td><input type=\"date\" name=\"datarilascio\" required></td</tr>");
    print("</table><br>");
    print("<input name=\"nomeis\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['nomeis']) . "\">");
    print("<input name=\"cognomeis\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['cognomeis']) . "\">");
    print("<input name=\"datais\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['datais']) . "\">");
    print("<input name=\"ntessera\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['ntessera']) . "\">");
    print("<input name=\"genitoreis\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['genitoreis']) . "\">");
    print("<input name=\"piscina\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['piscina']) . "\">");
    print("<input name=\"datainizio\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['datainizio']) . "\">");
    print ("<tr><td><input type=\"submit\" name=\"toinsert\" value=\"Insert\"></td></tr>");
    print ("</form>");

?>