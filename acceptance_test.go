// +build acceptance

package mongo

import (
	"fmt"
	"os"
	"testing"

	"github.com/ddspog/bdd"
)

const (
	// Const values to help tests legibility.
	testid01 = "000000007465737469643031"
	testid02 = "000000007465737469643032"
	testid03 = "000000007465737469643033"
	testid04 = "000000007465737469643034"
)

// TestMain setup the tests to run with real instance of MongoDB.
func TestMain(m *testing.M) {
	os.Setenv("MONGODB_URL", "mongodb://localhost:27017/test")

	if err := Connect(); err != nil {
		panic(err)
	}
	defer Disconnect()

	retCode := m.Run()
	defer os.Exit(retCode)
}

// Feature Manipulate data on MongoDB
// - As a developer,
// - I want to be able to connect and manipulate data from MongoDB,
// - So I can create an real application on this, using a DB.
func Test_Manipulate_data_on_MongoDB(t *testing.T) {
	given := bdd.Sentences().Given()

	given(t, fmt.Sprintf("a local database with a products collection that contains documents with ids: '%[1]s', '%[2]s', '%[3]s'", testid01, testid02, testid03), func(when bdd.When) {
		p := newProductHandle()
		defer p.Close()

		when("using p.RemoveAll()", func(it bdd.It) {
			removeInfo, errRemoveAll := p.RemoveAll()

			it("should return no errors", func(assert bdd.Assert) {
				assert.NoError(errRemoveAll)
			})

			it("should remove 3 documents", func(assert bdd.Assert) {
				assert.Equal(3, removeInfo.Removed)
			})

			it("should have p.Count() return 0", func(assert bdd.Assert) {
				n, errCount := p.Clean().Count()

				assert.NoError(errCount)
				assert.Equal(0, n)
			})

			it("should have p.FindAll() return nothing", func(assert bdd.Assert) {
				products, errFindAll := p.Clean().FindAll()

				assert.NoError(errFindAll)
				assert.Equal(0, len(products))
			})
		})

		newId := NewID()

		when(fmt.Sprintf("using p.Insert() to add documents with ids: '%[1]s', '%[2]s', '%[3]s', and a new one", testid01, testid02, testid03), func(it bdd.It) {
			errInsertDoc01 := p.SetDocument(&product{
				IDV: ObjectIdHex(testid01),
			}).Insert()
			errInsertDoc02 := p.SetDocument(&product{
				IDV: ObjectIdHex(testid02),
			}).Insert()
			errInsertDoc03 := p.SetDocument(&product{
				IDV: ObjectIdHex(testid03),
			}).Insert()
			errInsertDoc04 := p.SetDocument(&product{
				IDV: newId,
			}).Insert()

			it("should return no errors", func(assert bdd.Assert) {
				assert.NoError(errInsertDoc01)
				assert.NoError(errInsertDoc02)
				assert.NoError(errInsertDoc03)
				assert.NoError(errInsertDoc04)
			})

			it("should have p.Count() return 4", func(assert bdd.Assert) {
				n, errCount := p.Clean().Count()

				assert.NoError(errCount)
				assert.Equal(4, n)
			})

			it("should have p.Find() return all documents", func(assert bdd.Assert) {
				product01, errFindDoc01 := p.SetDocument(&product{
					IDV: ObjectIdHex(testid01),
				}).Find()

				assert.NoError(errFindDoc01)
				assert.Equal(testid01, product01.ID().Hex())

				product02, errFindDoc02 := p.SetDocument(&product{
					IDV: ObjectIdHex(testid02),
				}).Find()

				assert.NoError(errFindDoc02)
				assert.Equal(testid02, product02.ID().Hex())

				product03, errFindDoc03 := p.SetDocument(&product{
					IDV: ObjectIdHex(testid03),
				}).Find()

				assert.NoError(errFindDoc03)
				assert.Equal(testid03, product03.ID().Hex())

				newProduct, errFindNewDoc := p.SetDocument(&product{
					IDV: newId,
				}).Find()

				assert.NoError(errFindNewDoc)
				assert.Equal(newId.Hex(), newProduct.ID().Hex())
			})
		})

		t := NowInMilli()

		when(fmt.Sprintf("using p.Update() to change created_on doc with id '%[1]s' to '%[2]v'", newId.Hex(), t), func(it bdd.It) {
			errUpdate := p.SetDocument(&product{
				CreatedOnV: t,
			}).Update(newId)

			it("should return no errors", func(assert bdd.Assert) {
				assert.NoError(errUpdate)
			})

			it(fmt.Sprintf("should have p.Find() with id '%[1]s' return with new value", newId.Hex()), func(assert bdd.Assert) {
				product, errFind := p.SetDocument(&product{
					IDV: newId,
				}).Find()

				assert.NoError(errFind)
				assert.NotNil(product)
				assert.Equal(t, product.CreatedOn())
				assert.NotEqual(0, product.UpdatedOn())
			})

			it("should have p.Count() return 4", func(assert bdd.Assert) {
				n, errCount := p.Clean().Count()

				assert.NoError(errCount)
				assert.Equal(4, n)
			})
		})

		when(fmt.Sprintf("using p.Remove() to remove doc with id '%[1]s'", newId.Hex()), func(it bdd.It) {
			errRemove := p.Remove(newId)

			it("should return no errors", func(assert bdd.Assert) {
				assert.NoError(errRemove)
			})

			it(fmt.Sprintf("should have p.Find() with id '%[1]s' return nothing", newId.Hex()), func(assert bdd.Assert) {
				product, errFind := p.SetDocument(&product{
					IDV: newId,
				}).Find()

				assert.Error(errFind)
				assert.Contains(errFind.Error(), "not found")
				assert.Empty(product.ID().Hex())
			})

			it("should have p.Count() return 3", func(assert bdd.Assert) {
				n, errCount := p.Clean().Count()

				assert.NoError(errCount)
				assert.Equal(3, n)
			})
		})
	})
}

// Feature Read data on MongoDB
// - As a developer,
// - I want to be able to connect and retrieve data from MongoDB,
// - So I can use these functions on real applications.
func Test_Read_data_on_MongoDB(t *testing.T) {
	given, like, s := bdd.Sentences().All()

	given(t, fmt.Sprintf("a local database with a products collection that contains documents with ids: '%[1]s', '%[2]s', '%[3]s'", testid01, testid02, testid03), func(when bdd.When) {
		p := newProductHandle()
		defer p.Close()

		when("using p.Find() with '%[1]v' as document id'", func(it bdd.It, args ...interface{}) {
			product, errFind := p.SetDocument(&product{
				IDV: ObjectIdHex(args[0].(string)),
			}).Find()

			it("should run without errors", func(assert bdd.Assert) {
				assert.NoError(errFind)
			})

			it("should return a product with id '%[1]v'", func(assert bdd.Assert) {
				assert.Equal(args[0].(string), product.ID().Hex())
			})
		}, like(
			s(testid01), s(testid02), s(testid03),
		))

		when("using p.FindAll() with a empty Document", func(it bdd.It) {
			products, errFindAll := p.Clean().FindAll()

			it("should run without errors", func(assert bdd.Assert) {
				assert.NoError(errFindAll)
			})

			it("should return 3 documents", func(assert bdd.Assert) {
				assert.Equal(3, len(products))
			})

			if len(products) == 3 {
				it("should return the %[1]vth document with id '%[2]v'", func(assert bdd.Assert, args ...interface{}) {
					assert.Equal(args[1].(string), products[args[0].(int)-1].ID().Hex())
				}, like(
					s(1, testid01), s(2, testid02), s(3, testid03),
				))
			}
		})

		when("using p.FindAll() on a Document with id '%[1]v'", func(it bdd.It, args ...interface{}) {
			products, errFindAll := p.SetDocument(&product{
				IDV: ObjectIdHex(args[0].(string)),
			}).FindAll()

			it("should run without errors", func(assert bdd.Assert) {
				assert.NoError(errFindAll)
			})

			it("should return 1 documents", func(assert bdd.Assert) {
				assert.Equal(1, len(products))
			})

			if len(products) == 1 {
				it("should return the 1th document with id '%[1]v'", func(assert bdd.Assert) {
					assert.Equal(args[0].(string), products[0].ID().Hex())
				})
			}
		}, like(
			s(testid01), s(testid02), s(testid03),
		))
	})
}
