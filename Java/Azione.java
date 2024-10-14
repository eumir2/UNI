import java.time.LocalDate;

public class Azione extends Film{


//CONSTRUCTOR
    public Azione(String titolo, String tipo, LocalDate dN) {
        super(titolo, tipo, dN);
    }


    //METHODS
    @Override
    public double calcolaPenaleRitardo() {
        double tmp = (LocalDate.now().getDayOfMonth() - this.getDataNoleggio().getDayOfMonth()) - 7;
        if(tmp > 0)
            return tmp * 3;
        else
            throw new IllegalArgumentException("Film nolegiato da meno di una settimana");
    }
  
}
