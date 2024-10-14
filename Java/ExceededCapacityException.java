package PROVAESAME;

public class ExceededCapacityException extends Exception{
    
//CONSTRUCTORS
    public ExceededCapacityException(){
        super();
    }

    public ExceededCapacityException(String message){
        super(message);
    }
}
