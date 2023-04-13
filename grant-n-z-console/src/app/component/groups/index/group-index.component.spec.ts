import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupIndexComponent } from './group-index.component';

describe('IndexComponent', () => {
  let component: GroupIndexComponent;
  let fixture: ComponentFixture<GroupIndexComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GroupIndexComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupIndexComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
