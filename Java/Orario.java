import javax.lang.model.util.ElementScanner14;

public class Orario {
//OVERVIEW:che rappresenta un'orario con un valore delle ore tra 0 e 23 e dei minuti tra 0 e 59

    private int ore;
    private int minuti;
    private boolean formato;


    public static void main(String[] args) {
        Orario o;
        if(args.length == 0)
            o = new Orario();

        int minuti;
        int ore;
        boolean formato;

        if(args.length == 2){
            ore = Integer.parseInt(args[0]);
            minuti = Integer.parseInt(args[1]);
            try{
                o = new Orario(ore,minuti); 
            }catch(IllegalHoursException | IllegalMinutesException ex){
                System.out.println(ex.getMessage());
            }  
             
        }           

        if(args.length == 3){
            ore = Integer.parseInt(args[0]);
            minuti = Integer.parseInt(args[1]);    
            String tmp = args[2];
            if(tmp.equals("AM")) formato = false;
            else    formato = true;
            
            try{
                o = new Orario(ore,minuti,formato); 
            }catch(IllegalHoursException | IllegalMinutesException ex){
                System.out.println(ex.getMessage());
            }
        }





    }
   

//CONSTRUCTORS
    public Orario(int ore, int minuti, boolean formato) throws IllegalHoursException, IllegalMinutesException{
        
        check();
        this.minuti = minuti;
        this.formato = formato;
        this.ore = ore%12;
    }


    public Orario(int ore, int minuti) throws IllegalHoursException, IllegalMinutesException {

        check();

        this.ore = ore;
        this.minuti = minuti;
    }


    public Orario() {
        this.ore = 0;
        this.minuti = 0;
    }

//METHODS
    public void check()throws IllegalHoursException, IllegalMinutesException{
        if(this.ore < 0 || this.ore > 23){
            throw new IllegalHoursException("Valore ora invalido");
        }
        if(this.minuti < 0 || this.minuti > 59){
            throw new IllegalMinutesException("valore minuti errato");
        }
    }

    public int getOre() {
        if((Boolean)this.formato != null)
            return (ore+12);
        else
            return ore;
    }


    public int getMinuti() {
        return minuti;
    }

    public void avanza(int ore, int minuti){
        if(this.minuti + minuti >= 60){
            this.ore++;
            this.minuti = (this.minuti + minuti) % 60;
        }

        if(this.ore + ore >= 24){
            this.ore = (this.ore + ore) % 24;
        }
    }

    public String getOra24(){
        if((Boolean)this.formato != null){
            if(this.formato)
                return "Orario: "+(ore+12)+":"+minuti+" PM";
            else{
                return "Orario: "+ore+":"+minuti+" AM";
            }
        }
        return "Orario: "+ore+":"+minuti;
    }

    public String getOra12(){
        if((Boolean)this.formato != null){
            String tmp = formato ? " PM" : " AM";
            return "Orario: "+ore+" : "+minuti+" "+tmp; 
        }
        String tmp;
        if(this.ore > 12)
            tmp = "PM";
        else
            tmp = "AM";
        
        return "Orario: "+(ore%12)+" : "+minuti+" "+tmp;

    }
    


    public boolean repOk(){
        if((Integer)ore == null) return false;    
        if((Integer)minuti == null) return false;
        if((Boolean)formato == null) return false;

        if(this.ore < 0 || this.ore > 23)return false;
        if(this.minuti < 0 || this.minuti > 59)return false;
        
        return true;
    }

    @Override
    public boolean equals(Object obj) {
        if(obj != null) return false;
        if(!(obj instanceof Orario)) return false;

        Orario o = (Orario)obj;
        if((this.minuti != o.minuti) || (this.ore != o.ore) || (this.formato != o.formato))
            return false;
        
        return true;
    }
}