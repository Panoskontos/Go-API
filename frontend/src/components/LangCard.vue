<template>


      <transition-group enter-active-class="animate__animated animate__rollIn" leave-active-class="animate__animated animate__backOutLeft">
      
  <div v-for="lang in languages" :key="lang.id" class="card">
    <div class="face face1">
      <div class="content">
        <h2 class="python">
          <button class="vote_button vote" v-on:click="vote(lang.id);" style="-webkit-text-fill-color: #ffffff;">
            Vote
            </button>
            &nbsp;
          <button class="vote_button downvote" v-on:click="$emit('downvote',lang.id)" style="-webkit-text-fill-color: #ffffff;">
            Downvote
            </button>
            
          <i @click="$emit('removeid',lang.id)" class="bi bi-x-circle"></i>
          
           <i @click="$emit('editid',lang.id)" class="bi bi-gear-fill"></i>
          </h2>
        <br>
        <h2 class="python">{{lang.title}}</h2>
        <textarea v-model="newtext" @keyup.enter="changetext" class="edit_text invisible" name="" id="" cols="30" rows="10" :placeholder="lang.text"></textarea>
        <p class="python vis_text">{{lang.text}}</p>
      </div>
    </div>
    <div class="face face2">
      
      <h2>{{lang.votes}}</h2>
    </div>
  </div>
</transition-group>
</template>

<script>

export default {
     name: 'LangCard',
    data(){
    return {
      newtext: '',
    }},
     props: {languages: Array},
      methods: {
    //   longer way to do it
      vote(id){
          this.$emit('voteid',id)
      },
     
     


  }

}


</script>

<style>
    .bi {
        cursor:pointer;
        margin: 0 5px;
        -webkit-text-fill-color:rgb(255, 0, 255)
    }

    .edit_text {
        background-color: black;
        color:rgb(255, 255, 255);
        margin: 10px 0px;
        padding: 10px;
         border: 1px solid rgb(241, 135, 255);
         box-shadow: 1px 1px 1px rgb(249, 200, 255);
    }

    

</style>