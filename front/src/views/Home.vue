<template>
    <div>
        <h1 class="title is-1 has-text-centered">
            exonsunities
        </h1>
        <Card v-for="Community in communities" :key="Community._id" 
            :id="Community._id"
            :title="Community.title" :description="Community.description"/>
    </div>
</template>

<script>
import axios from 'axios'
import Card from '@/components/Card'

export default {
    name: 'Home',
    components: { Card },
    data() {
        return {
            communities: []
        }
    },
    beforeMount() {
        axios.get('http://localhost:8080/api/community/get/all').then(res => {
            this.communities = res.data;
        }).catch(err => {
            console.log(err);
        });
    }
}
</script>
