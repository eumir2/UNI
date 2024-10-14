import java.util.*;

public class CalcolatriceConMemoria{
/*  OVERVIEW: una calcolatrice sequenziale su numeri decimali, 
    utilizzando il risultato dell'operazione precedente 
    come il primo operando dell'operazione successiva.*/

    private double mem;


//CONSTRUCTORS
    public CalcolatriceConMemoria(){
    //MODIFIES: this
    //EFFECTS: inizializza this con memoria a 0
        this.mem = 0.0;
    }

    public CalcolatriceConMemoria(double mem){
    //MODIFIES: this
    //EFFECTS: inizializza this con memoria a mem
        this.mem = mem;
    }


    public double getMem(){
    //EFFECTS: ritorna il valore in memoria
        return this.mem;
    }

    public double add(double op2){
    //MODIFIES: this
    //EFFECTS: imposta il valore in memoria alla somma tra il valore corrente e op2 e lo ritorna
        mem = mem + op2;
        return mem;
    }

    public double sub(double op2){
    //MODIFIES: this
    //EFFECTS: imposta il valore in memoria alla differenza tra il valore corrente e op2 e lo ritorna  
        mem = mem - op2;
        return mem;
    }

    public double mul(double op2){
    //MODIFIES: this
    //EFFECTS: imposta il valore in memoria al prodotto tra il valore corrente e op2 e lo ritorna
    mem = mem * op2;
    return mem;
    }


    public double div(double op2) throws DivideByZeroException{
    //MODIFIES: this
    //EFFECTS: imposta il valore in memoria alla somma tra il valore corrente e op2 e lo ritorna, se op2 = 0 lancia DivideByZeroException
        if (op2 == 0){
            throw new DivideByZeroException("Non si può dividere per 0");
        } 

        mem = mem / op2;
        return mem;
    }

    public double operate(char operator, double op2) throws InputMismatchException, DivideByZeroException{
    //MODIFIES: this
    //EFFECTS: esegue la operazione definita in operator, se operator non è valido, lancia InputMismatchException
        switch (operator){
            case '+' : return add(op2);
            case '-' : return sub(op2);
            case '*' : return mul(op2);
            case '/' : return div(op2);
            default : throw new InputMismatchException("L'operatoe deve essere `+`, `-`, `*` e `/`");
        }
    }
}