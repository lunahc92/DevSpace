<template>
    <div>
        <v-switch v-model="autoDelete"
                  label="자동삭제"
        >
        </v-switch>
        <v-btn @click="showState()">show state</v-btn>
    </div>
</template>

<script>
    export default {
        name: 'DevSpaceAutoDeleter',
        props: ['autoDelete'],
        methods: {
            showState() {
                alert(this.autoDelete);
            }
        }
    }
</script>

<style>
    /* 대충 공대생 디자인... */
</style>