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
    $conn = pg_connect("host=localhost port=5432 dbname=Piscine user=postgres password=unimi");
    if (!$conn) {//caso connessione fallita
        echo 'Connessione al database fallita.';
        exit();
    }
    $query = "SELECT * FROM tipologia" ;   
   


    //print_R($_POST);
    $numvasche = (int)$_POST['nvasche'];
    //$row = pg_fetch_array($tipologie);
    //var_dump($numvasche);
    echo '<h1>Inserimento Vasche</h1>';

    print ("<form action=\"ins.php\" method=\"POST\">");
    echo '<table>';
    for($i = 1; $i <= $numvasche; $i++){
        $tipologie = pg_query($conn, $query);
        print("<tr><th>Vasca</th><td><input type=\"text\" name=\"codice[]\" value=\"". $i ."\" disabled></td></tr>");
        print ("<tr><th>Numero Corsie</th><td><input type=\"number\" name=\"ncors[]\"  min = \"1\" required></td></tr>");
        print ("<tr><th>Tipologia</th><td><select name=\"tipologia[]\" required>");
        //var_dump(sizeof($row));
        while($row = pg_fetch_array($tipologie)){
            echo("<option>". $row['tipo']."</option>");
        }
        
        print("</select>");
        echo "</td</tr>";
        print ("<tr><th>Inizio fruizione</th><td><input type=\"date\" name=\"iniziof[]\" required></td></tr>");
        print ("<tr><th>Fine fruizione</th><td><input type=\"date\" name=\"finef[]\" required></td></tr>");
    }


    print ("</table><br>");
    print("<input name=\"telc\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['telc']) . "\">");
    print("<input name=\"nome\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['nome']) . "\">");
    print("<input name=\"via\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['via']) . "\">");
    print("<input name=\"nciv\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['nciv']) . "\">");
    print("<input name=\"nvasche\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['nvasche']) . "\">");
    print("<input name=\"inizio\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['inizio']) . "\">");
    print("<input name=\"fine\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['fine']) . "\">");
    print("<input name=\"CF\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['CF']) . "\">");
    print("<input name=\"teler\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['telr']) . "\">");
    print("<input name=\"nomer\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['nomer']) . "\">");
    print("<input name=\"cognomer\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['cognomer']) . "\">");

    $diasettimana = array("Lunedi", "Martedi", "Mercoledi", "Giovedi", "Venerdi", "Sabato", "Domenica");
    foreach ($diasettimana as $giorno) {
        if(isset($_POST[$giorno])){
            print("<input name=".$giorno." type=\"hidden\" value=\"" . htmlspecialchars($_POST[$giorno]) . "\">");
        }
    }
    print("<input name=\"datar\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['datar']) . "\">");



    print ("<tr><td><input type=\"submit\" name=\"toinsert\" value=\"Insert\"></td></tr>");
    print ("</form>");
    print ($htmlint);
?>