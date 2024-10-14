<html>
    <head>
        <title>Inserimento e modifica</title>
        <style>
            table, th, td {
            text-align: left;
            border: 1px solid;
        }
        </style>
    </head>
    <body>
        <h1>Inserimento e modifica</h1>
        <div>
            <?php
            print("<form name =\"inserimento\" action =\"opt_basic.php\" method =\"POST\">
                Selezionare l'operazione da svolgere: 
                <select name = \"operazione\">
                    <option value = \"insert\">Inserimento</option>
                    <option value = \"update\">Modifica</option>
                </select>
                <br><br>
                Selezionare cosa si vuole modificare/inserire:
                <select name = \"tabella\">
                    <option value = \"piscina\">Struttura</option>
                    <option value = \"istruttore\">Personale</option>
                    <option value = \"corso\">Corso</option>
                    <option value = \"persona\">Iscritto</option>
                </select>
                <br>
                <input type = \"submit\">
            </form>");
            ?>
        </div>
    </body>


</html>