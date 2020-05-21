<template>
    <div class="root">
        <div class="alert alert-success"><i class="fas fa-check pr-2"></i> {{success}}</div>
        <div class="alert alert-info"><i class="fas fa-info pl-2 pr-2"></i> {{info}}</div>
        <div class="alert alert-danger"><i class="fas fa-exclamation pl-2 pr-2"></i> {{error}}</div>
        <div class="alert alert-warning"><i class="fas fa-exclamation pl-2 pr-2"></i> {{warning}}</div>

        <div class="form-group">
            <label for="success">Success</label>
            <input class="form-control" id="success" type="text" v-model="success">
        </div>

        <div class="form-group">
            <label for="info">Info</label>
            <input class="form-control" id="info" type="text" v-model="info">
        </div>

        <div class="form-group">
            <label for="error">Error</label>
            <input class="form-control" id="error" type="text" v-model="error">
        </div>

        <div class="form-group">
            <label for="warning">Warning</label>
            <input class="form-control" id="warning" type="text" v-model="warning">
        </div>
    </div>
</template>

<script type="module">
    export default {
        data() {
            return {
                success: 'Success',
                info: 'Info',
                error: 'Error',
                warning: 'Warning',
            }
        }
    };
</script>