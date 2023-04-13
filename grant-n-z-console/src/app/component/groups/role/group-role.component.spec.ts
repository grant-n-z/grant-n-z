import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupRoleComponent } from './group-role.component';

describe('GroupRoleComponent', () => {
  let component: GroupRoleComponent;
  let fixture: ComponentFixture<GroupRoleComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GroupRoleComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupRoleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
