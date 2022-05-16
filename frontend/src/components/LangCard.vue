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
          </h2>
        <br>
        <h2 class="python">{{lang.title}}</h2>
        <p class="python">{{lang.text}}</p>
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
  props: {languages: Array},
  methods: {
    //   longer way to do it
      vote(id){
          this.$emit('voteid',id)
      }
  }

}


</script>