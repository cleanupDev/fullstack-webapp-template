from dataclasses import dataclass


@dataclass
class User:
    id: int = None
    username: str = None
    password: str = None
    email: str = None
    first_name: str = None
    last_name: str = None
    created_at: str = None
    is_active: bool = True
    is_authenticated: bool = True
    
    def to_dict(self):
        return {'id': self.id, 'username': self.username, 'email': self.email, 'password': self.password, 'first_name': self.first_name, 'last_name': self.last_name}
    
    def get_id(self):
        return self.id
        

   