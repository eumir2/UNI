import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        Sfera sf = null;
        Cono c = null;

        System.out.println("Inserisci il raggio di una sfera");
        while(s.hasNext()){
            int raggioSfera = s.nextInt();
            sf = new Sfera("sfera", raggioSfera);

            System.out.println("Inserisci il raggio e l'altezza di un cono");
            int raggioCono = s.nextInt();
            int altezzaCono = s.nextInt();

            c = new Cono("cono", raggioCono, altezzaCono);
         }
         s.close();


         //System.out.println(sf.volume() + "    " + c.volume());
         
         int value = sf.compareTo(c);
         switch(value){
            case -1 : 
                System.out.println("La sfera ha valume minore del cono");
                break;
            case 0 :
                System.out.println("La sfera ha raggio uguale al cono");
                break;
            case 1 : 
                System.out.println("La sfera ha volume maggiore del cono");
                break;
         }


    }
}
