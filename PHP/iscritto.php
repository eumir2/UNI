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
function check($query, $conn){
    //var_dump($query);

    $result = pg_query($conn,$query);
    if (!$result) {
        echo "Si è verificato un errore.<br/>";
        echo pg_last_error($conn);
        exit();
    }
    return $result;
};

print($htmlint);
$conn = pg_connect("host=localhost port=5432 dbname=Piscine user=postgres password=unimi");
if (!$conn) {//caso connessione fallita
  echo 'Connessione al database fallita.';
  exit();
}

$tmp = isset($_POST['iscritto']) ? $_POST['iscritto'] : NULL;
$datiIscritto = explode(" ", $tmp);
//var_dump($datiIscritto);
$nome = substr($datiIscritto[0], 1);
$cognome = substr($datiIscritto[1], 0, -1);
//var_dump($nome);

//var_dump($cognome);

$nomepiscina = isset($_POST['piscina']) ? $_POST['piscina'] : NULL;
$numerotessera = isset($_POST['numerotessera']) ? $_POST['numerotessera'] : NULL;
//var_dump($piscina);

$query = "  SELECT *
            from tesseramento as t JOIN iscritto as i ON t.numerotessera = i.numerotessera 
            WHERE t.nomepiscina = '$nomepiscina'
            AND (i.nome = '$nome' AND i.cognome = '$cognome')
            ORDER BY t.inizio;";
$iscritti  = check($query, $conn);

if(pg_num_rows($iscritti) == 0){
    print("Non ci sono iscritti in questa piscina");
    echo ("<br><br>Se vuoi puoi <a href='visualizza.php'>riprovare</a>");
    exit();
}

//visualizza dati iscritto
print("<h1>Dati iscritto</h1>");
echo '<br><table>
    <tr>
    <th>Numero Tessera</td>
    <th>Nome</td>
    <th>Cognome</th>
    <th>Data di Nascita</th>
    <th>Genitore</th>
    <th>Inizio Iscrizione</th>
    <th>Fine Iscrizione</th>
    </tr>';
    while ($row = pg_fetch_array($iscritti)){
        print("<td> " .$row['numerotessera'] . "</td>");
        print("<td> " .$row['nome'] . "</td>");
        print("<td> " .$row['cognome'] . "</td>");
        print("<td> " .$row['datadinascita'] . "</td>");
        print("<td> " .$row['genitore'] . "</td>");
        print("<td> " .$row['inizio'] . "</td>");
        print("<td> " .$row['fine'] . "</td></tr>");
    }
    echo "</table>";
    echo ("<br><br>Operazione avvenuta con successo, <a href='visualizzazione.php'>RITORNA</a>");
?>

