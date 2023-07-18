class User:
    def __init__(self, **kwargs):
        self.id = kwargs.get('id')
        self.username = kwargs.get('username')
        self.email = kwargs.get('email')
        self.password = kwargs.get('password')
        self.first_name = kwargs.get('first_name')
        self.last_name = kwargs.get('last_name')
        self.is_authenticated = True
        self.is_active = True
        
    def __repr__(self):
        return {'id': self.id, 'username': self.username, 'email': self.email, 'password': self.password, 'first_name': self.first_name, 'last_name': self.last_name}
    
    def get_id(self):
        return self.id
        

   