//package Sottostringa;



/*Scrivere una programa che legga da riga di comando due stringhe. Se la seconda stringa è contenuta nella prima, scrivere a standard output <stringa2> è sottostringa di <stringa1>, altrimenti <stringa2> non è sottostringa di <stringa1>, come da esempi d'esecuzione.

A tal fine specificare, implementare e utilizzare la seguente procedura parziale:

public static boolean sottoStringa (String testo, String parola) che restituisce se true se parola è sottostringa di testo e false altrimenti. */
import java.util.Scanner;

public class Sottostringa {
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        String testo = s.next();
        String parola = s.next();
        s.close();

        //System.out.println(testo + "=>" + parola);

    

    }

    public static boolean sottoStringa (String testo, String parola){
        for (int i = 0; i < testo.length(); i++) {
            for(int j = 0; j < parola.length(); j++){
                if(testo.charAt(i) != parola.charAt(j)){
                    break;
                }
                i++;
                if j
            }
        }
    }
}
