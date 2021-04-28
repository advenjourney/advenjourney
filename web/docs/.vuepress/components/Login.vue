<template>
    <div>
        <div v-if="state === 'unauthorized'">
            <div>
                <h2>Login</h2>
                <label>
                    Username
                    <input type="text" v-model="loginUsername" />
                </label>
                </br>
                <label>
                    Password
                    <input type="password" v-model="loginPassword" />
                </label>
                <button @click="login">Login</button>
            </div>
            <div>
                <h2>Register</h2>
                <label>
                    Username
                    <input type="text" v-model="registerUsername" />
                </label>
                </br>
                <label>
                    Password
                    <input type="password" v-model="registerPassword" />
                </label>
                <button @click="register">Register</button>
            </div>
        </div>
        <div v-if="error" v-text="error" />
        <div v-if="state === 'authorized'">
            <p v-text="'Successfully logged in.'" />
        </div>
    </div>
</template>

<script>
export default {
    data: () => ({
        state: 'unauthorized',
        loginUsername: '',
        loginPassword: '',
        registerUsername: '',
        registerPassword: '',
        error: null,
        token: null,
    }),

    mounted() {
        this.apiUrl = this.$site.themeConfig.custom.url
        if (sessionStorage.getItem('token')) {
            console.log('Exists')
            this.state = 'authorized'
        }
    },

    methods: {
        login() {
            const storage = window.sessionStorage
            const query = `
            mutation {
                login(input: {username: "${this.loginUsername}", password: "${this.loginPassword}"})
            }
            `
            this.axios.post(this.apiUrl, { query })
                .then(res => {
                    if (res.data.data) {
                        this.state = 'authorized'
                        sessionStorage.setItem('token', res.data.data.login )

                        return
                    }
                    
                    this.error = res.data.errors[0].message
                }).catch(err => {
                    this.error = err
                })
        },

        register() {
            const query = `
            mutation {
                createUser(input: {username: "${this.registerUsername}", password: "${this.registerPassword}"})
            }
            `
            this.axios.post(this.apiUrl, { query })
                .then(res => {
                    if (res.data.data) {
                        this.state = 'authorized'
                        sessionStorage.setItem('token', res.data.data.createUser)

                        return
                    }
                    
                    this.error = res.data.errors[0].message
                }).catch(err => {
                    this.error = err
                })
        }
    }
}
</script>