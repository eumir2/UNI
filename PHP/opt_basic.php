<?php
session_start();
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

//print_R($_POST);
if (isset($_POST['tabella']) && isset($_POST['operazione'])) {//sono stati passati correttamente i dati
  $tbl = $_POST['tabella'];
  $op = $_POST['operazione'];
  $_SESSION['tbl'] = $tbl;
  //$_SESSION['op'] = $op;
  $conn = pg_connect("host=localhost port=5432 dbname=Piscine user=postgres password=unimi");
  if (!$conn) {//caso connessione fallita
    echo 'Connessione al database fallita.';
    exit();
  }
  else {//caso connessione riuscita
    //echo "Connessione riuscita."."<br/>";  
    $query = "SELECT * FROM tipologia" ;   
    $tipologie = pg_query($conn, $query);
    $query = "SELECT * FROM piscina"; 
    $piscine = pg_query($conn, $query);
    $query = "SELECT * FROM qualifica"; 
    $qualifiche = pg_query($conn, $query); 

    switch($op){  
        case 'insert':
            //l'inserimento richiede l'operazione preliminare di raccolta dei dati da inserire
            switch ($tbl) {
                case 'piscina':
                    print ($htmlint);
                    
                    echo '<h1>Inserimento struttura</h1>';
                    print ("<form action=\"vasche.php\" method=\"POST\">");
                    
                    echo '<table>';
                    //passo le informazioni oggetto dell'inserimento
                    print ("<tr><th>Telefono</th><td><input type=\"text\" name=\"telc\" maxlength = \"10\"required></td></tr>");
                    print ("<tr><th>Nome</th><td><input type=\"text\" name=\"nome\" required></td</tr>");
                    print ("<tr><th>Via</th><td><input type=\"text\" name=\"via\" required></td></tr>");
                    print ("<tr><th>Numero Civico</th><td><input type=\"number\" name=\"nciv\" required></td></tr>");
                    print ("<tr><th>Numero Vasche</th><td><input type=\"number\" min = \"1\" name=\"nvasche\"  required></td></tr>");
                    print ("<tr><th>Inizio apertura</th><td><input type=\"date\" name=\"inizio\" required></td></tr>");
                    print ("<tr><th>Fine apertura</th><td><input type=\"date\" name=\"fine\" required></td></tr>");
                    print ("<tr><th>Nome Responsabile</th><td><input type=\"text\" name=\"nomer\" required></td></tr>");
                    print ("<tr><th>Cognome Responsabile</th><td><input type=\"text\" name=\"cognomer\" required></td></tr>");
                    print ("<tr><th>Codice Fiscale Responsabile</th><td><input type=\"text\" name=\"CF\" maxlength = \"16\" required></td></tr>");
                    print ("<tr><th>Telefono Responsabile</th><td><input type=\"text\" name=\"telr\" maxlength = \"10\"required></td></tr>");
                    print ("<tr><th>Data di Nascita Responsabile</th><td><input type=\"date\" name=\"datar\" required></td></tr>");
                    
                    print ("<tr><th>reperibilità</th>");
                    $diasettimana = array("Lunedi", "Martedi", "Mercoledi", "Giovedi", "Venerdi", "Sabato", "Domenica");

                    echo '<td><ul>';
                    foreach ($diasettimana as $giorno) {
                        echo '<li><input type="checkbox" name="' . $giorno . '">' . $giorno . '</li>';
                    }
                    echo '</tr></td></ul>';
                        

                    print ("</table><br>");
                    print ("<tr><td><input type=\"submit\" name=\"toinsert\" value=\"Insert\"></td></tr>");
                    print ("</form>");
                break;
                case 'istruttore' :
                    print ($htmlint);
                    
                    echo '<h1>Inserimento istruttore</h1>';
                    print ("<form action=\"ins.php\" method=\"POST\">");
                    echo '<table>';
                    //passo le informazioni oggetto dell'inserimento
                    print ("<tr><th>Telefono</th><td><input type=\"text\" name=\"teli\" maxlength = \"10\"required></td></tr>");
                    print ("<tr><th>Nome</th><td><input type=\"text\" name=\"nomei\" required></td</tr>");
                    print ("<tr><th>Cognome</th><td><input type=\"text\" name=\"cognomei\" required></td</tr>");
                    print ("<tr><th>Data di nascita</th><td><input type=\"date\" name=\"datai\" required></td></tr>");
                    print ("<tr><th>Qualifiche</th>");

                    echo '<td><ul>';
                    while ($row = pg_fetch_array($qualifiche)){
                        echo("<li><input type = 'checkbox' name = '". $row['nome']. "'>". $row['nome'] ."</input></li>");
                    }
                    echo '</ul></td>';

                    print ("<tr><th>Nome Piscina</th><td><select name=\"nomep\" required>");
                    while ($row = pg_fetch_array($piscine)){
                        echo("<option>". $row['nome'] ."</option>");
                    }

                    print ("<tr><th>Tipo contratto</th><td><select name=\"contratto\" required>");
                    echo("<option>Indeterminato</option>");
                    echo("<option>Mensile</option>");
                    echo("<option>Bi-mestrale</option>");
                    echo("<option>Tri-mestrale</option>");
                    echo("<option>Se-mestrale</option>");
                    print ("<tr><th>Data inizio impiego</th><td><input type=\"date\" name=\"inizioi\" required></td></tr>");
                    print ("<tr><th>Gioni di ferie</th><td><input type=\"number\" min = \"0\" name=\"ferie\"  ></td></tr>");

                    //spiego nel documento


                    print ("</table><br>");
                    print ("<tr><td><input type=\"submit\" name=\"toinsert\" value=\"Insert\"></td></tr>");
                    print ("</form>");
                break;
                case 'corso' :
                    print ($htmlint);
                    
                    echo '<h1>Inserimento corso</h1>';
                    print ("<form action=\"ins.php\" method=\"POST\">");
                    echo '<table>';
                    //passo le informazioni oggetto dell'inserimento
                    print ("<tr><th>Nome Corso</th><td><input type=\"text\" name=\"nomec\" required></td></tr>");
                    print ("<tr><th>Nome Piscina</th><td><select name=\"nomep\" required>");
                    while ($row = pg_fetch_array($piscine)){
                        echo("<option>". $row['nome'] ."</option>");
                    }
                    echo "</td</tr>";

                    print ("<tr><th>Tipologia</th><td><select name=\"tipologiac\" required>");
                    while ($row = pg_fetch_array($tipologie)){
                        echo("<option>". $row['tipo'] ."</option>");
                    }
                    echo "</td</tr>";

                    print ("<tr><th>Livello</th><td><select name=\"livelloc\" required>");
                        echo("<option>". 1 ."</option>");
                        echo("<option>". 2 ."</option>");
                        echo("<option>". 3 ."</option>");
                    echo "</td</tr>";
                    print ("</table><br>");
                    print ("<tr><td><input type=\"submit\" name=\"toinsert\" value=\"Insert\"></td></tr>");
                    print ("</form>");
                break;
                case 'persona':
                    $tmp = "SELECT * FROM EdizioneCorso";
                    $edizionicorso = pg_query($conn, $tmp);

                    $tmp = "SELECT * FROM Piscina";
                    $piscine = pg_query($conn, $tmp);




                    print ($htmlint);
                    
                    echo '<h1>Inserimento iscritto</h1>';
                    print ("<form action=\"certificatomedico.php\" method=\"POST\">");
                    echo '<table>';
                    //passo le informazioni oggetto dell'inserimento
                    print ("<tr><th>Nome</th><td><input type=\"text\" name=\"nomeis\" required></td</tr>");
                    print ("<tr><th>Cognome</th><td><input type=\"text\" name=\"cognomeis\" required></td</tr>");
                    print ("<tr><th>NumeroTessera (10 cifre)</th><td><input type=\"text\" name=\"ntessera\"  minlength=\"10\" maxlength=\"10\" required></td></tr>");
                    print ("<tr><th>Data di nascita</th><td><input type=\"date\" name=\"datais\" required></td></tr>");
                    print ("<tr><th>Nome Genitore (opzionale)</th><td><input type=\"text\" name=\"genitoreis\"></td></tr>");

                    print ("<tr><th>Piscina</th><td><select name=\"piscina\" >");
                    while ($row = pg_fetch_array($piscine)){
                        echo("<option>". $row['nome'] ."</option>");
                    }
                    print ("<tr><th>Inizio Iscrizione</th><td><input type=\"date\" name=\"datainizio\" required></td></tr>");


                    print ("</table><br>");
                    print ("<tr><td><input type=\"submit\" name=\"toinsert\" value=\"Insert\"></td></tr>");
                    print ("</form>");
                break;
            }
            //  
        break;  
        case 'update':
            //l'operazione di aggiornamento richiede la scelta preliminare della tupla da aggiornare
            switch ($tbl) {
            case 'piscina':
                $query = "SELECT * FROM Piscina";
                $result = pg_query($conn, $query);
                if (!$result) {
                echo "Si è verificato un errore.<br/>";
                echo pg_last_error($conn);
                exit();
                }
                else {
                print ($htmlint);
                echo '<br>
            <table><tr>
            <th>Nome</td>
            <th>Via</th>
            <th>Numero Civico</th>
            <th>CF Responsabile</th>
            <th>Numero Vasche</th>
            <th>Inizio apertura</th>
            <th>Fine Apertura</th>
            </tr>';
                print ("<form action=\"upd.php\" method=\"POST\">");
                //passo le informazioni della tupla selezionata da aggiornare
                while ($row = pg_fetch_array($result)) {
                    echo "<tr>
                <td>" . $row['nome'] . "</td>
                <td>" . $row['via'] . "</td>        
                <td>" . $row['n_civ'] . "</td>
                <td>" . $row['cfresponsabile'] . "</td>        
                <td>" . $row['nvasche'] . "</td>       
                <td>" . $row['apertura'] . "</td>      
                <td>" . $row['chiusura'] . "</td>
                <td><input type=\"radio\" name=\"toupdate\" value='". $row['nome'] ."' required></td>
            </tr>";
                };
                echo '</table>';
                echo "<input type=\"submit\" name=\"update\" value=\"Update\">";
                echo "</form>";
                };
            break;
            case 'istruttore' :
                $query = "SELECT * FROM Istruttore";
                $result = pg_query($conn, $query);
                if (!$result) {
                echo "Si è verificato un errore.<br/>";
                echo pg_last_error($conn);
                exit();
                }
                else {
                print ($htmlint);
                echo '<br><table>
                        <tr>
                        <th>Nome</td>
                        <th>Cognome</th>
                        <th>Numero Badge</th>
                        <th>Data di nascita</th>
                        </tr>';
                print ("<form action=\"upd.php\" method=\"POST\">");
                //passo le informazioni della tupla selezionata da aggiornare
                while ($row = pg_fetch_array($result)) {
                    echo "<tr>
                    <td>" . $row['nome'] . "</td>
                    <td>" . $row['cognome'] . "</td>        
                    <td>" . $row['numerobadge'] . "</td>
                    <td>" . $row['datadinascita'] . "</td>             
                    <td><input type=\"radio\" name=\"toupdate\" value='" . $row['numerobadge'] . "' required></td>
                    </tr>";
                };
                echo '</table>';
                echo "<input type=\"submit\" name=\"update\" value=\"Update\">";
                echo "</form>";
                };
            break;
            case 'corso' :
                $query = "SELECT * FROM offre AS o JOIN edizionecorso AS e ON o.nomecorso = e.nomecorso ORDER BY O.nomepiscina";
                $result = pg_query($conn, $query);
                if (!$result) {
                echo "Si è verificato un errore.<br/>";
                echo pg_last_error($conn);
                exit();
                }
                else {
                print ($htmlint);
                echo '<br><table>
                        <tr>
                        <th>Nome Corso</td>
                        <th>Nome Piscina</th>
                        <th>Livello</th>
                        <th>Numero minimo partecipanti</th>
                        <th>Numero massimo partecipanti</th>
                        <th>Numero Lezioni</th>
                        <th>Costo</th>
                        </tr>';
                print ("<form action=\"upd.php\" method=\"POST\">");
                //passo le informazioni della tupla selezionata da aggiornare
                while ($row = pg_fetch_array($result)) {
                    echo "<tr>
                    <td>" . $row['nomecorso'] . "</td>
                    <td>" . $row['nomepiscina'] . "</td>        
                    <td>" . $row['livellocorso'] . "</td>
                    <td>" . $row['numerominpartecipanti'] . "</td>
                    <td>" . $row['numeromaxpartecipanti'] . "</td>
                    <td>" . $row['numerolezioni'] . "</td>
                    <td>" . $row['costo'] . "</td>              
                    <td><input type=\"radio\" name=\"toupdate\" value='" . $row['nomecorso'] . "-".  $row['nomepiscina'] . "-".  $row['livellocorso'] ."-". $row['codicecorso'] . "' required></td>
                    </tr>";
                };
                echo '</table>';
                echo "<input type=\"submit\" name=\"update\" value=\"Update\">";
                echo "</form>";
                };
            break;
            case 'persona' :
                $query = "SELECT * FROM iscritto as i JOIN tesseramento as t ON i.numerotessera = t.numerotessera";
                $result = pg_query($conn, $query);
                if (!$result) {
                echo "Si è verificato un errore.<br/>";
                echo pg_last_error($conn);
                exit();
                }
                else {
                print ($htmlint);
                echo '<br><table>
                        <tr>
                        <th>Numero Tessera</td>
                        <th>Nome</td>
                        <th>Cognome</th>
                        <th>Genitore</th>
                        <th>Data di Nascita</th>
                        <th>Nome Piscina</td>
                        <th>Fine Iscrizione</th>
                        </tr>';
                print ("<form action=\"upd.php\" method=\"POST\">");
                //passo le informazioni della tupla selezionata da aggiornare
                while ($row = pg_fetch_array($result)) {
                    echo "<tr>
                    <td>" . $row['numerotessera'] . "</td>
                    <td>" . $row['nome'] . "</td>
                    <td>" . $row['cognome'] . "</td>        
                    <td>" . $row['genitore'] . "</td>  
                    <td>" . $row['datadinascita'] . "</td>    
                    <td>" . $row['nomepiscina'] . "</td>
                    <td>" . $row['fine'] . "</td>              
                    <td><input type=\"radio\" name=\"toupdate\" value='" . $row['numerotessera'] . "' required></td>
                    </tr>";
                };
                echo '</table>';
                echo "<input type=\"submit\" name=\"update\" value=\"Update\">";
                echo "</form>";
                };
            break;
            };
        break;
    };
  };
}
else {//non sono stati passati correttamente i dati
  print ($htmlint);
  echo "Non risultano dati passati<br>";
  echo "Se vuoi puoi <a href='select_basic.php'>riprovare</a>";
}
?>
</BODY>
</HTML>