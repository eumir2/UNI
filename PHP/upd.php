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

//funzionce che esegue le query
function check($query, $conn){
  //var_dump($query);

  $result = pg_query($conn,$query);
  if (!$result) {
      echo "Si è verificato un errore.<br/>";
      echo pg_last_error($conn);
      echo ("<br><br>Se vuoi puoi <a href='inserimento_e_modifica.php'>riprovare</a>");
      exit();
  }
  return $result;
}


print($htmlint);
if (isset($_POST['toupdate']) && isset($_SESSION['tbl'])) {//sono stati passati correttamente i dati
    $conn = pg_connect("host=localhost port=5432 dbname=Piscine user=postgres password=unimi");
    if (!$conn) {
      echo 'Connessione al database fallita.';
      exit();
    }else{
        $tbl = $_SESSION['tbl'];
        switch($tbl){
            case 'piscina' :
                //estrapolo i dati della piscina
                $nome = isset($_POST['toupdate']) ? $_POST['toupdate'] : NULL;
                $query = "SELECT * FROM piscina WHERE piscina.nome = '$nome'";
                //var_dump($query);
                $result = check($query, $conn);
                print("<h1>Modifica Piscina</h1>");

                print("<table><tr>
                <th>Nome</td>
                <th>Via</th>
                <th>Numero Civico</th>
                <th>Responsabile</th>
                <th>Numero Vasche</th>
                <th>Inizio apertura</th>
                <th>Fine Apertura</th>
                </tr>");
                print ("<form action=\"upd2.php\" method=\"POST\">");
                //passo le informazioni della tupla selezionata da aggiornare
                while ($row = pg_fetch_array($result)) {
                    echo "<tr>";
                    print("<td><input type=\"text\" name=\"nome\" readonly value=\"". $row['nome']  ."\"></td>");
                    print("<td><input type=\"text\" name=\"via\"   readonly value=\"". $row['via']  ."\"></td>");
                    print("<td><input type=\"number\" name=\"nciv\"  value=\"". $row['n_civ']  ."\"></td>");
                    print("<td><input type=\"text\" name=\"cfres\" readonly value=\"". $row['cfresponsabile']  ."\"></td>");
                    print("<td><input type=\"number\" name=\"nvasche\"  value=\"". $row['nvasche']  ."\"></td>");
                    print("<td><input type=\"date\" name=\"inizio\"  value=\"". $row['apertura']  ."\"></td>");
                    print("<td><input type=\"date\" name=\"fine\"  value=\"". $row['chiusura']  ."\"></td>");
   
                };
                echo '</tr></table><br>';
                echo "<input type=\"submit\" name=\"update\" value=\"Update\">";
                echo "</form>";
              break;
              case 'istruttore' : 
                //estrapolo i dati  dell'istruttore
                $numerobadge = isset($_POST['toupdate']) ? $_POST['toupdate'] : NULL;
                $query = "SELECT * FROM istruttore WHERE istruttore.numerobadge = '$numerobadge'";
                //var_dump($query);

                $result = check($query, $conn);
                print("<h1>Modifica Personale</h1>");

                echo '<br><table>
                <tr>
                <th>Numero Badge</th>
                <th>Nome</td>
                <th>Cognome</th>
                <th>Data di nascita</th>
                </tr>';
                
                print ("<form action=\"upd2.php\" method=\"POST\">");
                //passo le informazioni della tupla selezionata da aggiornare
                while ($row = pg_fetch_array($result)) {
                    echo '<tr>';
                    print("<td><input type=\"number\" name=\"nbadge\" readonly value=\"". $row['numerobadge']  ."\"></td>");
                    print("<td><input type=\"text\" name=\"nomei\"  value=\"". $row['nome']  ."\"></td>");
                    print("<td><input type=\"text\" name=\"cognomei\"  value=\"". $row['cognome']  ."\"></td>");
                    print("<td><input type=\"date\" name=\"datai\"  value=\"". $row['datadinascita']  ."\"></td>");   
                    echo '</tr>';
                };
                echo '</table><br>';
                echo "<input type=\"submit\" name=\"update\" value=\"Update\">";
                echo "</form>";
              break;
              case  'corso' :
                //estrapolo i dati del corso
                $dati = isset($_POST['toupdate']) ? $_POST['toupdate'] : NULL;
                if($dati){
                  $p = explode('-',$dati);
                }
                //var_dump($p);
                $query = "SELECT * FROM offre AS o JOIN edizionecorso AS e ON o.nomecorso = e.nomecorso 
                WHERE o.nomecorso = '$p[0]'
                  AND o.nomepiscina = '$p[1]'
                  AND e.livellocorso = '$p[2]'
                  AND e.codicecorso = '$p[3]';";
                //var_dump($query);
                $result = check($query, $conn);

                print("<h1>Modifica Corso</h1>");
                echo '<br><table>
                <tr>
                <th>Codice Corso</td>
                <th>Nome Corso</td>
                <th>Nome Piscina</th>
                <th>Livello</th>
                <th>Numero minimo partecipanti</th>
                <th>Numero massimo partecipanti</th>
                <th>NumeroLezioni</th>
                <th>Costo</th>          
                </tr>';
                print ("<form action=\"upd2.php\" method=\"POST\">");
                while ($row = pg_fetch_array($result)){
                  //print_r($row);
                  echo '<tr>';
                  print("<td><input type=\"text\" name=\"codicecorso\" readonly value=\"". $row['codicecorso']  ."\"></td>");
                  print("<td><input type=\"text\" name=\"ncorso\" readonly value=\"". $row['nomecorso']  ."\"></td>");
                  print("<td><input type=\"text\" name=\"npiscina\" readonly value=\"". $row['nomepiscina']  ."\"></td>");
                  print("<td><input type=\"text\" name=\"livelloc\" readonly value=\"". $row['livellocorso']  ."\"></td>");
                  print("<td><input type=\"text\" name=\"minpar\"  value=\"". $row['numerominpartecipanti']  ."\"></td>");
                  print("<td><input type=\"text\" name=\"maxpar\"  value=\"". $row['numeromaxpartecipanti']  ."\"></td>");
                  print("<td><input type=\"text\" name=\"nlezioni\"  value=\"". $row['numerolezioni']  ."\"></td>");
                  print("<td><input type=\"text\" name=\"costo\"  value=\"". $row['costo']  ."\"></td>");
                };
                
                echo '</tr></table><br>';
                echo "<input type=\"submit\" name=\"update\" value=\"Update\">";
                echo "</form>";
              break;
              case 'persona' :
                //estrapolo i dati della persona
                $numerotessera = isset($_POST['toupdate']) ? $_POST['toupdate'] : NULL;
                $query = "SELECT * FROM iscritto as i JOIN tesseramento as t ON i.numerotessera = t.numerotessera WHERE i.numerotessera = '$numerotessera'";
                $result = pg_query($conn, $query);
                if (!$result) {
                echo "Si è verificato un errore.<br/>";
                echo pg_last_error($conn);
                exit();
                }

                $result = check($query, $conn);

                print("<h1>Modifica iscritto</h1>");
                echo '<br><table>
                <tr>
                <th>Numero Tessera</td>
                <th>Nome</td>
                <th>Cognome</th>
                <th>Genitore</th>
                <th>Data di Nascita</th>
                <th>Nome Piscina</td>
                <th>Inizio Iscrizione</th>
                <th>Fine Iscrizione</th>
                </tr>';

                print ("<form action=\"upd2.php\" method=\"POST\">");
                while ($row = pg_fetch_array($result)){
                  print("<td><input type=\"text\" name=\"ntessera\" readonly value=\"". $row['numerotessera']  ."\"></td>");
                  print("<td><input type=\"text\" name=\"nomeis\"  value=\"". $row['nome']  ."\"></td>");
                  print("<td><input type=\"text\" name=\"cognomeis\"  value=\"". $row['cognome']  ."\"></td>");
                  print("<td><input type=\"text\" name=\"genitoreis\"  value=\"". $row['genitore']  ."\"></td>");
                  print("<td><input type=\"date\" name=\"datais\"  value=\"". $row['datadinascita']  ."\"></td>");
                  print("<td><input type=\"text\" name=\"nomepiscina\" readonly value=\"". $row['nomepiscina']  ."\"></td>");
                  print("<td><input type=\"date\" name=\"inizio\"  readonly value=\"". $row['inizio']  ."\"></td>");
                  print("<td><input type=\"date\" name=\"fine\"  value=\"". $row['fine']  ."\"></td>");
                }
                echo "</table><br>";
                echo "<input type=\"submit\" name=\"update\" value=\"Update\">";
                echo "</form>";
              break;
        };
      }
    }
?>