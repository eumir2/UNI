import java.util.Scanner;

public class SvolgiCalcolo {
    public static void main(String[] args) {

        CalcolatriceConMemoria c;

       /*try {
        double tmp = Double.parseDouble(args[0]);
         c = new CalcolatriceConMemoria(tmp);
       } catch (ArrayIndexOutOfBoundsException e) {
         c = new CalcolatriceConMemoria();
       }*/
       
       if (args.length > 0) c = new CalcolatriceConMemoria(Double.parseDouble(args[0]));
        else c = new CalcolatriceConMemoria();
        
       Scanner s = new Scanner(System.in);

       while(true){
            char op = s.next().charAt(0);

            if(op == '=')   break;

            double op2 = s.nextDouble();

            try {
                System.out.println(c.operate(op, op2));    
            } catch (DivideByZeroException e) {
                System.out.println(e.getMessage());
            }
         }

    }
}
