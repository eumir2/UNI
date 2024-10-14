import java.util.ArrayList;
import java.util.Scanner;

import javax.sound.midi.Soundbank;

public class Rubrica {
//OVERVIEW: classe che testa la classe Contatto

    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        ArrayList<Contatto> contatti = new ArrayList<Contatto>();

        while(s.hasNext()){
            String[] parametri = s.nextLine().split(" ");
            //System.out.println(parametri);
            if(parametri[2].equals(null))
                parametri[2] = " ";

            if(parametri[3].equals(null))
                parametri[3] = " ";
                
            Contatto tmp = new Contatto(parametri[0], parametri[1], parametri[2], parametri[3]);
            
            contatti.add(tmp);
        }

        ArrayList<Contatto> contattiUguali = verificaUguali(contatti);
            if(contattiUguali.size() == 0)
                System.out.println("Non ci sono contatti uguali");
            else{
                System.out.println("Contatti uguali: ");
                for (Contatto c : contattiUguali) {
                    System.out.println(c.toString());
                }
            }
    }

    public static ArrayList<Contatto> verificaUguali(ArrayList<Contatto> contatti) {
        ArrayList<Contatto> tmp = new ArrayList<Contatto>();
        for (int i = 0; i < contatti.size()-1; i++) {
            for (int j = i+1; j < contatti.size(); j++) {
                Contatto c1 = contatti.get(i);
                Contatto c2 = contatti.get(j);
                
                if((c1 != c2) && (c1.equals(c2)))
                    tmp.add(c1);
            }            
        }
        return tmp;
    }
    
}
