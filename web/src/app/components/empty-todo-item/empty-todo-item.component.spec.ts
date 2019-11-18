import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EmptyTodoItemComponent } from './empty-todo-item.component';

describe('EmptyTodoItemComponent', () => {
  let component: EmptyTodoItemComponent;
  let fixture: ComponentFixture<EmptyTodoItemComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EmptyTodoItemComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EmptyTodoItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
