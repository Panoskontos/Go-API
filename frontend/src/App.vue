<template>
  <div>
    <h1 class="myhead" style="color: #ffffff;">Vote your Favourite Language</h1>

    <!-- <div>
      <div v-for="lang in languages" :key="lang.id">
        {{lang.title}}
        {{lang.votes}}
        
      </div>
    </div> -->
    <div class="container">
 <!-- animation -->

  <LangCard @voteid="vote" @downvote="downvote" :languages="languages"/>


  <!-- <div class="card">
    <div class="face face1">
      <div class="content">
        <span class="stars"></span>
        <h2 class="python">Python</h2>
        <p class="python">Python is an interpreted, high-level and general-purpose programming language.</p>
      </div>
    </div>
    <div class="face face2">
      <h2>02</h2>
    </div>
  </div> -->

  <!-- <div class="card">
    <div class="face face1">
      <div class="content">
        <span class="stars"></span>
        <h2 class="cSharp">C#</h2>
        <p class="cSharp">C# is a general-purpose, multi-paradigm programming language encompassing static typing, strong typing, lexically scoped and component-oriented programming disciplines.</p>
      </div>
    </div>
    <div class="face face2">
      <h2>03</h2>
    </div>
  </div> -->
</div>
  </div>
</template>

<script>
import axios from 'axios';
import LangCard from './components/LangCard.vue'

const baseURL = 'http://localhost:3000/posts'
const voteURL = "http://localhost:3000/vote_post?id="
const downvoteURL = "http://localhost:3000/downvote_post?id="


export default {
  name: 'App',
  components: {
    LangCard,
  },
  data(){
    return {
      languages: [],
  }
},
 async created(){
    try {
      const res = await axios.get(baseURL);
      console.log(res.data)
      this.languages = res.data;
    } catch(e) {
      console.error(e)
    }
  },
   methods: {
     async vote(id){
      const res = await axios.patch(voteURL + id);
      console.log(res.data)
      this.languages[id-1].votes +=1
      var btn = document.querySelectorAll('.vote')[id-1]
      btn.disabled = true;
     },

     async downvote(id){
      const res = await axios.patch(downvoteURL + id);
      console.log(res.data)
      if (this.languages[id-1].votes >0) {
          this.languages[id-1].votes -=1
      }
      var btn = document.querySelectorAll('.downvote')[id-1]
      btn.disabled = true;
      
     }
   }
}
</script>


<!-- css -->
<style>
@import url("https://fonts.googleapis.com/css2?family=Righteous&display=swap");
body {
	margin: 0;
	padding: 0;
	display: flex;
	justify-content: center;
	align-items: center;
	font-family: "Righteous", cursive;
	min-height: 100vh;
	background-color: #a9c9ff;
	background-image: linear-gradient(180deg, #a9c9ff 0%, #ffbbec 100%);
}

.vote_button {
  
  font-size: 16px;
  padding: 3px 15px;
  border-radius: 10px;
  cursor:pointer;
  background: linear-gradient(217deg, rgba(255, 255, 255, 0.8), rgba(255,0,0,0) 70.71%),
            linear-gradient(127deg, rgba(0, 238, 255, 0.8), rgba(0,255,0,0) 70.71%),
            linear-gradient(336deg, rgba(255, 0, 230, 0.8), rgba(0,0,255,0) 70.71%);
}

.myhead {
  display: flex;
  justify-content: center;
  margin:10px;
}

 body .container {
	max-width: 100vw;
	display: grid;
	grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
	grid-gap: 35px;
	margin: 0 auto;
	padding: 40px 0;
}
 body .container .card {
	position: relative;
	width: 300px;
	height: 400px;
	margin: 0 auto;
	background: #000;
	border-radius: 15px;
	box-shadow: 0 15px 60px rgba(0, 0, 0, 0.5);
}
 body .container .card .face {
	position: absolute;
	bottom: 0;
	left: 0;
	width: 100%;
	height: 100%;
	display: flex;
	justify-content: center;
	align-items: center;
}
 body .container .card .face.face1 {
	box-sizing: border-box;
	padding: 20px;
}
 body .container .card .face.face1 h2 {
	margin: 0;
	padding: 0;
}
 body .container .card .face.face1 .java {
	background-color: #fffc00;
	-webkit-background-clip: text;
	-webkit-text-fill-color: transparent;
}
 body .container .card .face.face1 .python {
	background-color: #ffffff;
	-webkit-background-clip: text;
	-webkit-text-fill-color: transparent;
}
 body .container .card .face.face1 .cSharp {
	background-color: #fc00ff;
	-webkit-background-clip: text;
	-webkit-text-fill-color: transparent;
}
 body .container .card .face.face2 {
	transition: 0.5s;
}
 body .container .card .face.face2 h2 {
	margin: 0;
	padding: 0;
	font-size: 10em;
	color: #fff;
	transition: 0.5s;
	text-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
	z-index: 10;
}
 body .container .card:hover .face.face2 {
	height: 60px;
}
 body .container .card:hover .face.face2 h2 {
	font-size: 2em;
}
 body .container .card .face.face2 {
	background-image: linear-gradient(40deg, #c806c1 0%, #10f1f1 45%, #c500bf 100%);
	border-radius: 15px;
}
 /* body .container .card:nth-child(2) .face.face2 {
	background-image: linear-gradient(40deg, #000000 0%, #000000 45%, #000000 100%);
	border-radius: 15px;
}
 body .container .card:nth-child(3) .face.face2 {
	background-image: linear-gradient(40deg, #000000 0%, #000000 45%, #000000 100%);
	border-radius: 15px;
}
 body .container .card:nth-child(4) .face.face2 {
	background-image: linear-gradient(40deg, #000000 0%, #000000 45%, #000000 100%);
	border-radius: 15px;
}
 body .container .card:nth-child(5) .face.face2 {
	background-image: linear-gradient(40deg, #000000 0%, #000000 45%, #000000 100%);
	border-radius: 15px; */
/* } */
 /* body .container .card .face.face2 {
	background-image: linear-gradient(40deg,  #c806c1 0%, #10f1f1 45%, #c500bf 100%);
	border-radius: 15px;
} */
 
</style>
