import java.util.ArrayList;
import java.util.Iterator;
import java.util.Scanner;



public class Pets {



//MAIN
    public static void main(String[] args) {
        ArrayList<Persona> persone = new ArrayList<Persona>();

        Scanner s = new Scanner(System.in);
        System.out.println("Inserisci righe nel formato `nomePersona nomeAnimale tipoAnimale` (CTRL+D per terminare)");

        while(s.hasNext()){
            Persona p = new Persona(s.next());
            Boolean flag = false;
            
            //System.out.println(persone.contains(p));
    
            String nomePet = s.next();
            String tipoPet = s.next();
            Pet pet;

            if(tipoPet.equals("Cane"))
                pet = new Cane(nomePet);
            else if(tipoPet.equals("Gatto"))
                pet = new Gatto(nomePet);
            else
                pet = new Cavia(nomePet);
            
            if(persone.size() < 1){
                p.addAnimale(pet);
                persone.add(p);
            }else{
                for (int i = 0; i < persone.size(); i++) {
                    if(persone.get(i).equals(p)){
                        persone.get(i).addAnimale(pet);
                        break;
                    }
                    else{
                        flag = true;
                        break;
                    }
                }
                if(flag){
                    p.addAnimale(pet);
                    persone.add(p);
                }

            }
        }
        System.out.println(persone.size());
        s.close();

        for (Persona persona : persone) {
            System.out.println("Coro degli animali di " + persona.getNome());
            Iterator<Pet> i = persona.iterator();
            while(i.hasNext()){
                Pet tmp = i.next();
                System.out.print(tmp.getNome() + " dice ");
                tmp.verso();
                System.out.println();
            }
            System.out.println();

            
        }
    
    }
    
}
