import { async, ComponentFixture, TestBed } from "@angular/core/testing";

import { BodyComponent } from "./body.component";
import { ItemService } from "src/app/services/item.service";
import { TodoCardComponent } from "../todo-card/todo-card.component";
import { HttpClient } from "@angular/common/http";

describe("BodyComponent", () => {
  let component: BodyComponent;
  let fixture: ComponentFixture<BodyComponent>;

  beforeEach(async(() => {
    const spy = jasmine.createSpyObj("ItemService", [
      "getItems",
      "getAllItems",
      "http"
    ]);
    TestBed.configureTestingModule({
      declarations: [BodyComponent, TodoCardComponent],
      providers: [
        {
          provide: ItemService,
          useValue: spy
        }
      ]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BodyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it("should create", () => {
    expect(component).toBeTruthy();
  });
});
