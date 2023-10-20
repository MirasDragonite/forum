const indexOf=(a,b,c)=>{
let count=c
let len= a.length 

if (count=== undefined || count<=0){
    count=0
}
while(count<len){

if(a[count]===b){
return count
}
    count++
}

return -1
}

const lastIndexOf=(a,b)=>{
   
    let len= a.length-1
    
 
    while(len>=0){
    
    if(a[len]===b){
    return len
    }
      len--
    }
    
    return -1
    }

    const includes=(a,b)=>{
   
        let len= a.length-1
        
     
        while(len>=0){
        
        if(a[len]===b){
        return true
        }
          len--
        }
        
        return false
        }

